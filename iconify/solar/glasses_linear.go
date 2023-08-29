package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GlassesLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="18" cy="16" r="4"/><circle cx="6" cy="16" r="4"/><path stroke-linecap="round" d="m14 16.214l-.656-.234a3.999 3.999 0 0 0-2.688 0l-.656.234M2 16l.763-8.395c.115-1.264.173-1.896.543-2.363c.37-.467.972-.668 2.176-1.07L6 4m16 12l-.763-8.395c-.115-1.264-.172-1.896-.542-2.363c-.37-.467-.973-.668-2.177-1.07L18 4"/></g>`),
		g.Group(children),
	)
}