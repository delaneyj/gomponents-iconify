package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Opensource(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M15.5 1.125C7.278 1.125.59 7.815.59 16.035c0 6.263 3.88 11.635 9.36 13.84l3.64-9.076a5.131 5.131 0 1 1 3.818-.001l3.64 9.075c5.48-2.206 9.36-7.578 9.36-13.84c.002-8.22-6.687-14.91-14.908-14.91z"/>`),
		g.Group(children),
	)
}