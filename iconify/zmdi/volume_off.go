package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M288 192q0 6-1 13l-52-52v-47q24 12 38.5 35t14.5 51zm53 0q0-50-30-89.5T235 49V5q64 15 106.5 67T384 192q0 47-22 89l-32-33q11-27 11-56zM27 0l165 165l192 192l-27 27l-44-44q-35 29-78 39v-44q25-8 48-25l-91-91v144L85 256H0V128h101L0 27zm165 21v90l-45-45z"/>`),
		g.Group(children),
	)
}