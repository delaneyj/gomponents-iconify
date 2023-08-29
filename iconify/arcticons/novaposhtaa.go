package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Novaposhtaa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m5.5 24l7.488-7.488v14.976L5.5 24Zm37 0l-7.488-7.488v14.976L42.5 24Zm-22.024 3.524h7.048v7.488h3.964L24 42.5l-7.488-7.488h3.964v-7.488Zm0-7.048h7.048v-7.488h3.964L24 5.5l-7.488 7.488h3.964v7.488Z"/>`),
		g.Group(children),
	)
}