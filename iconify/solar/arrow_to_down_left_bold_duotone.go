package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowToDownLeftBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M17.53 13.97a.75.75 0 0 1 0 1.06l-5 5a.75.75 0 0 1-1.06 0l-5-5a.75.75 0 1 1 1.06-1.06L12 18.44l4.47-4.47a.75.75 0 0 1 1.06 0Z" clip-rule="evenodd"/><path d="M12.75 9.5c0-.953-.28-2.367-1.141-3.563C10.719 4.702 9.244 3.75 7 3.75a.75.75 0 0 0 0 1.5c1.756 0 2.78.715 3.391 1.563c.639.887.859 1.974.859 2.687v8.19l.75.75l.75-.75V9.5Zm-.552 10.724Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}