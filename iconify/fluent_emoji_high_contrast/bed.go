package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M15 20H7v-3.5a.5.5 0 0 0 .5.5h7a.5.5 0 0 0 .5-.5v-.75c0-.946-.731-1.5-1.687-1.735a2.363 2.363 0 0 1-.123-.033c-.448-.137-3.287-.985-4.44-.982A1.75 1.75 0 0 0 7 14.75V12a3 3 0 1 0-6 0v15a2 2 0 0 0 2 2h1v1.5a.5.5 0 0 0 .5.5h2a.5.5 0 0 0 .5-.5V29h18v1.5a.5.5 0 0 0 .5.5h2a.5.5 0 0 0 .5-.5V29h1a2 2 0 0 0 2-2v-7.504a3.5 3.5 0 0 0-3.5-3.5h-11a1.5 1.5 0 0 0-1.5 1.5V20ZM4 11a1 1 0 0 1 1 1v10h10v.5a.5.5 0 0 0 .5.5H29v4H3V12a1 1 0 0 1 1-1Z"/>`),
		g.Group(children),
	)
}