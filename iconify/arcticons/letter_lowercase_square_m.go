package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterLowercaseSquareM(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.879 22.332a5.56 5.56 0 0 1 5.561-5.56h0a5.56 5.56 0 0 1 5.56 5.56v8.897M12.879 16.771v14.458M24 22.332a5.56 5.56 0 0 1 5.56-5.56h0a5.56 5.56 0 0 1 5.561 5.56v8.897"/>`),
		g.Group(children),
	)
}