package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M33.55 8.2a28.1 28.1 0 0 0-31.11.08a1 1 0 1 0 1.12 1.66a26.11 26.11 0 0 1 28.89-.07a1 1 0 0 0 1.1-1.67Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M18.05 10.72A20.74 20.74 0 0 0 6.23 14.4A1 1 0 0 0 7.36 16a18.85 18.85 0 0 1 21.28 0a1 1 0 0 0 1.12-1.65a20.75 20.75 0 0 0-11.71-3.63Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="M18.05 17.9a13.51 13.51 0 0 0-8 2.64a1 1 0 0 0 1.18 1.61a11.56 11.56 0 0 1 13.62-.08A1 1 0 1 0 26 20.46a13.52 13.52 0 0 0-7.95-2.56Z" class="clr-i-outline clr-i-outline-path-3"/><path fill="currentColor" d="M18 24.42a4 4 0 1 0 4 4a4 4 0 0 0-4-4Zm0 6a2 2 0 1 1 2-2a2 2 0 0 1-2 2Z" class="clr-i-outline clr-i-outline-path-4"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}