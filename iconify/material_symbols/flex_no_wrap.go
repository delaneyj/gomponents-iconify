package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlexNoWrap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 17V7h6v10H1Zm8 0V7h6v10H9Zm8 0V7h6v10h-6ZM3 15h2V9H3v6Zm16 0h2V9h-2v6Z"/>`),
		g.Group(children),
	)
}