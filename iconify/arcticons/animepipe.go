package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Animepipe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.003 35.727V19.75h3.674v-3.517H38.38v3.9h3.597l.02 15.594M8.733 37.458h30.454m-26.062.515V43m21.757-5.027V43M24 11.233l.004 3.271m-3.472-6.546l.004 3.27M17.065 5l.004 3.271m10.398-.313l.004 3.27M30.935 5l.004 3.271"/>`),
		g.Group(children),
	)
}