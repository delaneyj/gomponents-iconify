package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sharethis(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M164 232v5l165 82q22-19 52-19q34 0 57.5 23.5T462 381t-23.5 57.5T381 462t-57.5-23.5T300 381v-5l-165-82q-22 19-52 19q-34 0-57.5-23.5T2 232t23.5-57.5T83 151q30 0 52 19l165-82v-5q0-34 23.5-57.5T381 2t57.5 23.5T462 83t-23.5 57.5T381 164q-30 0-52-19l-165 82v5z"/>`),
		g.Group(children),
	)
}