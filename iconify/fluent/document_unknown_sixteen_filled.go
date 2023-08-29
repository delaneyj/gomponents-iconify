package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentUnknownSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M6 1.002h3v3.5a1.5 1.5 0 0 0 1.5 1.5H14v6.996a2 2 0 0 1-2 2H8.666A5.5 5.5 0 0 0 4 5.207V3.002a2 2 0 0 1 2-2z" fill="currentColor"/><path d="M10.5 5.003h3.497l-3.989-4.001H10v3.5a.5.5 0 0 0 .5.5z" fill="currentColor"/><path d="M10 10.5a4.5 4.5 0 1 1-9 0a4.5 4.5 0 0 1 9 0zm-4.5 1.88a.625.625 0 1 0 0 1.25a.625.625 0 0 0 0-1.25zm0-4.877c-1.048 0-1.864.818-1.853 1.955a.5.5 0 0 0 1-.01c-.006-.579.36-.945.853-.945c.472 0 .853.392.853.95c0 .202-.07.315-.36.544l-.277.215C5.21 10.616 5 10.929 5 11.5a.5.5 0 0 0 .992.09l.011-.156c.017-.148.1-.254.346-.448l.277-.215c.513-.41.727-.732.727-1.318c0-1.104-.822-1.95-1.853-1.95z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}