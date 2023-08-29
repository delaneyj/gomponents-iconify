package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M304 160c0-70.7-57.3-128-128-128h-32C73.3 32 16 89.3 16 160c0 34.6 13.7 66 36 89c-31.5 23.3-52 60.8-52 103c0 70.7 57.3 128 128 128h64c70.7 0 128-57.3 128-128c0-42.2-20.5-79.7-52-103c22.3-23 36-54.4 36-89zM176.1 288H192c35.3 0 64 28.7 64 64s-28.7 64-64 64h-64c-35.3 0-64-28.7-64-64s28.7-64 64-64h48.1zm0-64H144c-35.3 0-64-28.7-64-64s28.7-64 64-64h32c35.3 0 64 28.7 64 64s-28.6 64-64 64z"/>`),
		g.Group(children),
	)
}