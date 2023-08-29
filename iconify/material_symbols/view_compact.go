package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewCompact(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 19.75V4.25v15.5Zm-20-11v-4.5h4.25v4.5H2Zm5.25 0v-4.5h4.25v4.5H7.25Zm5.25 0v-4.5h4.25v4.5H12.5Zm5.25 0v-4.5H22v4.5h-4.25Zm0 5.5v-4.5H22v4.5h-4.25Zm-5.25 0v-4.5h4.25v4.5H12.5Zm-5.25 0v-4.5h4.25v4.5H7.25Zm-5.25 0v-4.5h4.25v4.5H2Zm15.75 5.5v-4.5H22v4.5h-4.25Zm-5.25 0v-4.5h4.25v4.5H12.5Zm-5.25 0v-4.5h4.25v4.5H7.25Zm-5.25 0v-4.5h4.25v4.5H2Z"/>`),
		g.Group(children),
	)
}