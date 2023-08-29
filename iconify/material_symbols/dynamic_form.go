package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DynamicForm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 11q-.825 0-1.413-.588T2 9V6q0-.825.588-1.413T4 4h9v7H4Zm0 9q-.825 0-1.413-.588T2 18v-3q0-.825.588-1.413T4 13h11v7H4Zm13 0v-9h-2V4h7l-2 5h2l-5 11ZM4.75 17.25h1.5v-1.5h-1.5v1.5Zm0-9h1.5v-1.5h-1.5v1.5Z"/>`),
		g.Group(children),
	)
}