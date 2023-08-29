package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M6 34a1 1 0 0 1-1-1V3a1 1 0 0 1 2 0v30a1 1 0 0 1-1 1Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M30.55 3.82a1 1 0 0 0-1 0a14.9 14.9 0 0 1-6.13 1.16a13.11 13.11 0 0 1-5.18-1.49a12.78 12.78 0 0 0-5-1.45A10.86 10.86 0 0 0 9 2.85v2.23A8.8 8.8 0 0 1 13.25 4a11.22 11.22 0 0 1 4.2 1.28a14.84 14.84 0 0 0 6 1.66A18.75 18.75 0 0 0 29 6.12v12.83a16.16 16.16 0 0 1-5.58.93a13.11 13.11 0 0 1-5.18-1.49a12.78 12.78 0 0 0-5-1.45a10.86 10.86 0 0 0-4.24.85V20a8.8 8.8 0 0 1 4.25-1.08a11.22 11.22 0 0 1 4.2 1.28a14.84 14.84 0 0 0 6 1.66a16.79 16.79 0 0 0 7-1.37a1 1 0 0 0 .55-.89V4.67a1 1 0 0 0-.45-.85Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}