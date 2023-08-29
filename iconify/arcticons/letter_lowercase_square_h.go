package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterLowercaseSquareH(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Zm-22.061 7.379v22.242"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.44 25.946a5.56 5.56 0 0 1 5.56-5.56h0a5.56 5.56 0 0 1 5.56 5.56v9.175"/>`),
		g.Group(children),
	)
}