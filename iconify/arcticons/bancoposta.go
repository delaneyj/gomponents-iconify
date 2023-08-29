package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bancoposta(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.056 31.5v-15h4.91c2.777 0 5.027 2.255 5.027 5.038s-2.25 5.037-5.027 5.037h-4.91M18.194 24a3.75 3.75 0 1 1 0 7.5h-6.187v-15h6.187a3.75 3.75 0 1 1 0 7.5h0Zm0 0h-6.187"/>`),
		g.Group(children),
	)
}