package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HourglassStart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.304 22a11.077 11.077 0 0 0-5.645-9.971L14.604 12c3.427-1.929 5.704-5.543 5.704-9.689c0-.109-.002-.218-.005-.327V2a.923.923 0 0 0 .899-1.003V1a.923.923 0 0 0-.898-1H.903a.923.923 0 0 0-.901 1.003V1a.923.923 0 0 0 .9 1h.001a11.077 11.077 0 0 0 5.645 9.971l.055.029C3.176 13.929.899 17.543.899 21.689c0 .109.002.218.005.327V22a1.007 1.007 0 0 0-.005 2h19.503a.923.923 0 0 0 .9-1.003V23a1 1 0 0 0-.999-1zm-17.7 0c0-5 3.5-9 8-9s8 4 8 9z"/>`),
		g.Group(children),
	)
}