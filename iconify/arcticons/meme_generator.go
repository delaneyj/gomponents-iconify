package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MemeGenerator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.48 5.5a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33.04a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.5 26.736v-3.367a2.104 2.104 0 0 1 2.104-2.105h0a2.104 2.104 0 0 1 2.105 2.105v3.367"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.709 23.369a2.104 2.104 0 0 1 2.104-2.105h0a2.104 2.104 0 0 1 2.104 2.105v3.367m5.807-1.01a2.103 2.103 0 0 1-1.829 1.062h0a2.104 2.104 0 0 1-2.104-2.104v-1.368a2.104 2.104 0 0 1 2.104-2.104h0A2.104 2.104 0 0 1 23 23.316V24h-4.208M25 26.736v-3.367a2.104 2.104 0 0 1 2.105-2.105h0a2.104 2.104 0 0 1 2.104 2.105v3.367m.001-3.367a2.104 2.104 0 0 1 2.104-2.105h0a2.104 2.104 0 0 1 2.104 2.105v3.367m5.806-1.01a2.104 2.104 0 0 1-1.828 1.062h0a2.104 2.104 0 0 1-2.105-2.104v-1.368a2.104 2.104 0 0 1 2.105-2.104h0a2.104 2.104 0 0 1 2.104 2.104V24h-4.209"/>`),
		g.Group(children),
	)
}