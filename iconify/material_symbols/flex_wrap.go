package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlexWrap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 22v-9h6v9H1Zm8 0v-9h6v9H9Zm8 0v-9h6v9h-6Zm-6-2h2v-5h-2v5ZM1 11V2h6v9H1Zm8 0V2h6v9H9Zm8 0V2h6v9h-6ZM3 9h2V4H3v5Zm16 0h2V4h-2v5Z"/>`),
		g.Group(children),
	)
}