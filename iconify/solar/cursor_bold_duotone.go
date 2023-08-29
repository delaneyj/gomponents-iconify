package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CursorBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="m11.433 16.464l1.203-1.202l2.626-2.626l1.202-1.203c1.232-1.23 1.847-1.846 1.702-2.508c-.146-.662-.963-.963-2.596-1.565l-5.45-2.007C6.861 4.152 5.231 3.55 4.391 4.39c-.84.84-.24 2.47.962 5.73l2.007 5.45c.602 1.633.903 2.45 1.565 2.596c.662.145 1.277-.47 2.508-1.702Z" clip-rule="evenodd"/><path d="m12.636 15.262l3.938 3.938c.408.408.612.612.84.706c.302.126.643.126.946 0c.228-.094.432-.298.84-.706c.407-.408.611-.612.706-.84a1.238 1.238 0 0 0 0-.946c-.095-.228-.299-.432-.706-.84l-3.939-3.938l-2.625 2.626Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}