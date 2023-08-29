package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatRoundCallBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12c0 1.6.376 3.112 1.043 4.453c.178.356.237.763.134 1.148l-.595 2.226a1.3 1.3 0 0 0 1.591 1.591l2.226-.595a1.634 1.634 0 0 1 1.149.133A9.958 9.958 0 0 0 12 22Z" opacity=".5"/><path fill-rule="evenodd" d="M12 6.775a.825.825 0 0 0-.825.825v8.8a.825.825 0 0 0 1.65 0V7.6A.825.825 0 0 0 12 6.775Zm-4.4 2.2a.825.825 0 0 0-.825.825v4.4a.825.825 0 0 0 1.65 0V9.8a.825.825 0 0 0-.825-.825Z" clip-rule="evenodd"/><path d="M15.575 9.8a.825.825 0 0 1 1.65 0v4.4a.825.825 0 0 1-1.65 0V9.8Z"/></g>`),
		g.Group(children),
	)
}