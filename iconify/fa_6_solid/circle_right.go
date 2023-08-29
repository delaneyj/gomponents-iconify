package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 256a256 256 0 1 0 512 0a256 256 0 1 0-512 0zm294.6-120.9l99.9 107.1c3.5 3.8 5.5 8.7 5.5 13.8s-2 10.1-5.5 13.8l-99.9 107.1c-4.2 4.5-10.1 7.1-16.3 7.1c-12.3 0-22.3-10-22.3-22.3V304h-96c-17.7 0-32-14.3-32-32v-32c0-17.7 14.3-32 32-32h96v-57.7c0-12.3 10-22.3 22.3-22.3c6.2 0 12.1 2.6 16.3 7.1z"/>`),
		g.Group(children),
	)
}