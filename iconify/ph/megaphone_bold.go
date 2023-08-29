package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MegaphoneBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M244 120a52.06 52.06 0 0 0-52-52h-39.68c-3.44-.21-52.6-4-99.46-43.3A20 20 0 0 0 20 40v160a19.8 19.8 0 0 0 11.54 18.12a19.86 19.86 0 0 0 21.32-2.81A192.92 192.92 0 0 1 136 174.47v26.2a20 20 0 0 0 8.9 16.64a11.35 11.35 0 0 0 1.39.8l14.44 7.06A20 20 0 0 0 190.37 213l11.09-41.82A52.07 52.07 0 0 0 244 120ZM44 191.63V48.4c36.17 28.07 72.17 38.1 92 41.66V150c-19.83 3.52-55.83 13.55-92 41.63Zm124.39 10.57l-8.39-4.1V172h16.4ZM192 148h-32V92h32a28 28 0 1 1 0 56Z"/>`),
		g.Group(children),
	)
}