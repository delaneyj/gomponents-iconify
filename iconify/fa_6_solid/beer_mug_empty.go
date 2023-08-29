package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BeerMugEmpty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M32 64c0-17.7 14.3-32 32-32h288c17.7 0 32 14.3 32 32v32h51.2c42.4 0 76.8 34.4 76.8 76.8v102.1c0 30.4-17.9 57.9-45.6 70.2L384 381.7V416c0 35.3-28.7 64-64 64H96c-35.3 0-64-28.7-64-64V64zm352 247.6l56.4-25.1c4.6-2.1 7.6-6.6 7.6-11.7v-102c0-7.1-5.7-12.8-12.8-12.8H384v151.6zM160 144c0-8.8-7.2-16-16-16s-16 7.2-16 16v224c0 8.8 7.2 16 16 16s16-7.2 16-16V144zm64 0c0-8.8-7.2-16-16-16s-16 7.2-16 16v224c0 8.8 7.2 16 16 16s16-7.2 16-16V144zm64 0c0-8.8-7.2-16-16-16s-16 7.2-16 16v224c0 8.8 7.2 16 16 16s16-7.2 16-16V144z"/>`),
		g.Group(children),
	)
}