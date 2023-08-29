package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WiFi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor"><path d="M8.646 17A4.489 4.489 0 0 1 12 15.5c1.333 0 2.53.58 3.354 1.5h.265l2.841-3.966A9.466 9.466 0 0 0 12 10.5a9.466 9.466 0 0 0-6.46 2.534L8.381 17h.265Z"/><path d="M20.218 10.58L23 6.5v-.25C20.09 4.051 16.5 2.5 12 2.5S3.91 4.05 1 6.25v.25l2.782 4.08A12.452 12.452 0 0 1 12 7.5c3.146 0 6.02 1.162 8.218 3.08ZM10.5 20a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Z"/></g>`),
		g.Group(children),
	)
}