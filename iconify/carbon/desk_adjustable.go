package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeskAdjustable(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M17 27.172V22h-2v5.172l-2.586-2.586L11 26l5 5l5-5l-1.414-1.414L17 27.172z"/><path fill="currentColor" d="M30 12H2v8h2v10h2V20h20v10h2V20h2zm-2 6H4v-4h24zM15 4.828V10h2V4.828l2.586 2.586L21 6l-5-5l-5 5l1.414 1.414L15 4.828z"/>`),
		g.Group(children),
	)
}