package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zohomail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 31.25v9.5a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-20.5l-18-15h-1l-18 15v6l18 9.54h1l11-5.92v-7.62a1.11 1.11 0 0 0-1-1h-21a1.12 1.12 0 0 0-1 1c-.05.5 0 2.92 0 2.92"/>`),
		g.Group(children),
	)
}