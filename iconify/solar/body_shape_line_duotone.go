package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BodyShapeLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M20 2s-2 4.688-2 8.571c0 1.244.426 2.284 1 3.32c.66 1.193 1.517 2.38 2.146 3.863c.499 1.178.854 2.543.854 4.246M4 2s2 4.688 2 8.571c0 1.244-.426 2.284-1 3.32c-.66 1.193-1.517 2.38-2.146 3.863A10.605 10.605 0 0 0 2 22"/><path d="M6 13h12M6 11h12" opacity=".5"/><path d="M12 22c.5-1.5 3-4.5 9-4.5M12 22c-.5-1.5-3-4.5-9-4.5"/></g>`),
		g.Group(children),
	)
}