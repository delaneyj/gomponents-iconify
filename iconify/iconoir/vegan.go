package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vegan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M15 11.063C12.53 13.65 10.059 20 10.059 20S6.529 11.062 3 9"/><path d="m20.496 5.577l.426 4.424c.276 2.87-1.875 5.425-4.745 5.702c-2.816.27-5.367-1.788-5.638-4.604a5.122 5.122 0 0 1 4.608-5.59l4.716-.454a.58.58 0 0 1 .633.522Z"/></g>`),
		g.Group(children),
	)
}