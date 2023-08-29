package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CursorLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m12.636 15.262l-1.203 1.202c-1.23 1.232-1.846 1.847-2.508 1.702c-.662-.146-.963-.963-1.565-2.596l-2.007-5.45C4.152 6.861 3.55 5.231 4.39 4.391c.84-.84 2.47-.24 5.73.962l5.45 2.007c1.633.602 2.45.903 2.596 1.565c.145.662-.47 1.277-1.702 2.508l-1.202 1.203"/><path d="m12.636 15.262l3.938 3.938c.408.408.612.612.84.706c.303.126.643.126.947 0c.227-.094.431-.298.839-.706s.611-.612.706-.84a1.238 1.238 0 0 0 0-.946c-.095-.228-.298-.432-.706-.84l-3.938-3.938" opacity=".5"/></g>`),
		g.Group(children),
	)
}