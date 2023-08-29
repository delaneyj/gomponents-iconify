package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Parimatchru(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.561 4.5h20.337a1.263 1.263 0 0 1 1.159 1.794l-4.173 14.14a2.531 2.531 0 0 1-2.218 1.795H7.329a1.263 1.263 0 0 1-1.159-1.795l4.173-14.14A2.53 2.53 0 0 1 12.561 4.5Zm7.773 21.271h20.337a1.263 1.263 0 0 1 1.159 1.795l-4.173 14.14a2.53 2.53 0 0 1-2.218 1.794H15.102a1.263 1.263 0 0 1-1.159-1.794l4.173-14.14a2.531 2.531 0 0 1 2.218-1.795Z"/>`),
		g.Group(children),
	)
}