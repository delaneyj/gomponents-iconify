package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Difference(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.5 11h2V9h2V7h-2V5h-2v2h-2v2h2v2Zm-2 4h6v-2h-6v2ZM8 19q-.825 0-1.413-.588T6 17V3q0-.825.588-1.413T8 1h7l6 6v10q0 .825-.588 1.413T19 19H8Zm-4 4q-.825 0-1.413-.588T2 21V7h2v14h11v2H4Z"/>`),
		g.Group(children),
	)
}