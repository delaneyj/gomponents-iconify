package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stopwatch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 4h4a.5.5 0 0 0 0-1h-4a.5.5 0 0 0 0 1zm7.993 4.713l1.36-1.36a.5.5 0 0 0-.707-.707l-1.36 1.36C15.876 6.762 14.03 6 12 6s-3.876.762-5.287 2.007l-1.36-1.36a.5.5 0 0 0-.707.707l1.36 1.36A7.96 7.96 0 0 0 4 14a8 8 0 0 0 8 8a8.01 8.01 0 0 0 8-8a7.96 7.96 0 0 0-2.007-5.287zM12 21A7 7 0 0 1 7.037 9.065c.005-.005.012-.006.017-.011c.005-.005.006-.012.01-.017A6.977 6.977 0 0 1 12 7a6.986 6.986 0 0 1 4.943 2.049l.008.008A6.986 6.986 0 0 1 19 14a7 7 0 0 1-7 7zm.5-8.408V10.5a.5.5 0 0 0-1 0v2.092A1.496 1.496 0 0 0 12 15.5a1.496 1.496 0 0 0 .5-2.908z"/>`),
		g.Group(children),
	)
}