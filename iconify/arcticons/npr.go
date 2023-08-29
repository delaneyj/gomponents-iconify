package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Npr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.641 21.425v5.152m0-3.244a1.91 1.91 0 0 1 1.907-1.907h0m-16.344-.124v7.633m0-4.389a1.91 1.91 0 0 0 1.908 1.907h0a1.91 1.91 0 0 0 1.907-1.907v-1.24a1.91 1.91 0 0 0-1.907-1.907h0a1.91 1.91 0 0 0-1.908 1.907m-13.42 3.271v-5.152m3.815 5.152V23.43a1.91 1.91 0 0 0-1.907-1.907h0a1.91 1.91 0 0 0-1.907 1.907m8.726-5.934H4.5v13.008h13.008l.003-13.008Zm13.008 0H17.51v13.008h13.008V17.496Zm-.027 0v13.008H43.5V17.495l-13.008.001Z"/>`),
		g.Group(children),
	)
}