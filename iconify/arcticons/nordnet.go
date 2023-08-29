package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nordnet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.15 17.35h10.1v13.3l11.65-13.3h5.6l-11.65 13.3h-10.1v-13.3L10.1 30.65H4.5Zm0 0"/>`),
		g.Group(children),
	)
}