package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsMergeAltV(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 1.503v2h-5v4.172l1.829-1.829l1.414 1.415L12 11.503L7.757 7.261l1.415-1.415L11 7.675V3.503H6v-2h12Zm0 18.994v2H6v-2h5v-4.172l-1.828 1.829l-1.415-1.415L12 12.497l4.243 4.242l-1.415 1.415L13 16.325v4.172h5Z"/>`),
		g.Group(children),
	)
}