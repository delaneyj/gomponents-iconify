package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnjoinThreeD(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.5 7L9 12h6l-2.5 5m8.339 3.84h-3.536m3.536 0v-3.537m0 3.536L17 17M2.768 2.768h3.535m-3.535 0v3.535m0-3.535l3.839 3.839"/>`),
		g.Group(children),
	)
}