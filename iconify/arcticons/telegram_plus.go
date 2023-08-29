package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TelegramPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 4.5A16.362 16.362 0 0 0 7.637 20.863A23.842 23.842 0 0 0 24 43.5v-6.275h0A16.363 16.363 0 0 0 24 4.5Zm0 21.745V15.48m-5.382 5.383h10.764"/>`),
		g.Group(children),
	)
}