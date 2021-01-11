package ing

import (
	"context"
	"fmt"
	"net/url"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"time"

	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/chromedp"
)

const (
	instagramHost = "instagram.com"
	magicNumber   = 20
)

// OpenLink opens the given link in a new tab in
// your default browser.
//
// Should fail if no Chrome presented on executor.
func OpenLink(raw string) error {
	ic := &gGetter{
		// c: &http.Client{
		// 	Transport: http.DefaultTransport,
		// 	Timeout:   time.Second * 5,
		// },
		rawURL: raw,
	}

	i, err := ic.getRawG()
	if err != nil {
		return fmt.Errorf("failed to get raw link: %w", err)
	}

	if _, err = url.Parse(i); err != nil {
		return fmt.Errorf("got unparseable url from the client: %s", i)
	}

	return performOpen(runtime.GOOS, i)
}

var cdnRe = regexp.MustCompile(`.*content="(.*(?:fbcdn\.net|cdninstagram\.com).*)".*`)

// TODO: remove usage of chromedp, it's too slow and won't work w/o chrome
func (s *gGetter) getRawG() (string, error) {
	initU, err := url.Parse(s.rawURL)
	if err != nil {
		return "", fmt.Errorf("failed to parse raw string '%s': %w", s.rawURL, err)
	}

	if !strings.Contains(initU.Host, instagramHost) {
		return "", fmt.Errorf("the initial was not an '%s': %s", initU.Host, instagramHost)
	}

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.NoFirstRun,
		chromedp.NoSandbox,
		chromedp.NoDefaultBrowserCheck,
		chromedp.Flag("headless", true),
		chromedp.Flag("ignore-certificate-errors", true),
	)
	ectx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	cctx, cancel := chromedp.NewContext(ectx)
	defer cancel()
	ctx, cancel := context.WithTimeout(cctx, time.Second*5)
	defer cancel()

	var res string
	if err = chromedp.Run(ctx,
		chromedp.Navigate(initU.String()),
		chromedp.ActionFunc(func(ctx context.Context) error {
			node, rerr := dom.GetDocument().WithDepth(1).Do(ctx)
			if rerr != nil {
				return rerr
			}
			res, err = dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)
			return err
		}),
	); err != nil {
		return "", fmt.Errorf("failed to run cdt: %w", err)
	}

	res = strings.Replace(res, "&amp;", "&", magicNumber)
	m := cdnRe.FindStringSubmatch(res)
	if len(m) != 2 {
		return "", fmt.Errorf("corrupted regexp match after getting body of '%s' and has %d matches", initU.String(), len(m))
	}

	return m[1], nil
}

func performOpen(goos, s string) error {
	var cmdA string
	switch runtime.GOOS {
	default:
		return fmt.Errorf("don't know what to do with '%s' platform", runtime.GOOS)
	case "darwin":
		cmdA = "open"
	case "windows":
		// https://stackoverflow.com/a/49115945/1561149
		cmdA = "rundll32 url.dll,FileProtocolHandler"
	case "linux":
		cmdA = "xdg-open"
	}

	args := []string{cmdA, s}
	cmd := exec.Command(args[0], args[1:]...)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to perform '%s': %w", cmd.String(), err)
	}
	return nil
}
