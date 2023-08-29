package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Doswimmer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.24 27.394h14.666c7.96 0 6.331-9.693.18-9.723c-.402-7.116-11.322-8.673-13.44-2.077c-7.227.11-8.371 11.8-1.406 11.8Zm8.207 5.59l-3.87-1.823l-2.483 3.93m1.062-1.2l2.52 1.173M5.5 35.227h37"/><circle cx="15.475" cy="32.486" r="1.349" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}