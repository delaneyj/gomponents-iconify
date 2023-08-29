package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AerialTramway(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m27.846 2.012l-6 .938l-12.5 1.95l-5.5.859a1 1 0 1 0 .308 1.976l5.5-.859l1.854-.29L15 6.043V10h-4a4 4 0 0 0-4 4v14a4 4 0 0 0 4 4h10a4 4 0 0 0 4-4V14a4 4 0 0 0-4-4h-4V5.731l4.433-.693l.721-.112l6-.938a1 1 0 1 0-.308-1.976ZM21 30H11a2 2 0 0 1-2-2h14a2 2 0 0 1-2 2Zm2-16h-5.07a.9.9 0 0 0-.93.861v7.278a.9.9 0 0 0 .93.861H23v3H9v-3h5.07a.9.9 0 0 0 .93-.861v-7.278a.9.9 0 0 0-.93-.861H9a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2Z"/>`),
		g.Group(children),
	)
}