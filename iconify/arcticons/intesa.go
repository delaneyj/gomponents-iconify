package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Intesa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2ZM8.614 16.158h30.772"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.614 31.842V23.45a4.386 4.386 0 0 1 8.772 0v8.392m-19.772 0V23.45a4.386 4.386 0 0 1 8.772 0v8.392m13.228 0V23.45a4.386 4.386 0 0 1 8.772 0v8.392"/>`),
		g.Group(children),
	)
}