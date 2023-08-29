package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Motoaudio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.5 34.717h4.333m4.334-21.434H16.5M20.833 24h4.334M29.5 13.283h4.333m4.334 21.434H42.5M5.667 38.5v-29m8.666 29v-29M23 38.5v-29m8.667 29v-29m8.666 29v-29"/>`),
		g.Group(children),
	)
}