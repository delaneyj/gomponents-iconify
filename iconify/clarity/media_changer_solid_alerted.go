package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MediaChangerSolidAlerted(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<g id="clarityMediaChangerSolidAlerted0" fill="currentColor"><path d="M22.23 15.4a3.68 3.68 0 0 1-3.18-5.51L22.45 4H6a2 2 0 0 0-2 2v24a2 2 0 0 0 2 2h1.88v1.57a1 1 0 0 0 2 0V32h16v1.57a1 1 0 0 0 2 0V32H30a2 2 0 0 0 2-2V15.4ZM17 28H8.81v-2H17Zm0-4H8.81v-2H17Zm0-4H8.81v-2H17Zm0-4H8.81v-2H17Zm0-4H8.81v-2H17Zm5 12h-2v-2h2Zm0-4h-2v-2h2Zm4 4h-2v-2h2Zm0-4h-2v-2h2Z"/><path d="m26.85 1.14l-5.72 9.91a1.27 1.27 0 0 0 1.1 1.95h11.45a1.27 1.27 0 0 0 1.1-1.91l-5.72-9.95a1.28 1.28 0 0 0-2.21 0Z"/></g>`),
		g.Group(children),
	)
}