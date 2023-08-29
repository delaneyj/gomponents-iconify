package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderFocusOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M43 23v-9a2 2 0 0 0-2-2H24l-5-6H7a2 2 0 0 0-2 2v32a2 2 0 0 0 2 2h15"/><path d="m35 27l2.243 4.913l5.365.615l-3.979 3.651l1.073 5.293L35 38.816l-4.702 2.656l1.073-5.293l-3.98-3.651l5.366-.615L35 27Z"/></g>`),
		g.Group(children),
	)
}