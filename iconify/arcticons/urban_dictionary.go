package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UrbanDictionary(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.513 23.575a4.5 4.5 0 0 0-4.5-4.5h0a4.5 4.5 0 0 0-4.5 4.5V26.5a4.5 4.5 0 0 0 4.5 4.5h0a4.5 4.5 0 0 0 4.5-4.5m1.8 4.5a1.8 1.8 0 0 1-1.8-1.8V13M22.041 26.5a4.5 4.5 0 0 1-4.5 4.5h0a4.5 4.5 0 0 1-4.5-4.5v-7.425M23.841 31a1.8 1.8 0 0 1-1.8-1.8V19.075m-9 0h-1.354m10.354 0h-1.354M34.513 13h-1.354"/>`),
		g.Group(children),
	)
}