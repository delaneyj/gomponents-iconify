package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PorcelainVase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M304 30c-32 64-22.35 180.063 0 208c42.406 53.007 80 80 80 128c0 32-16 80-48 112v16H176v-16c-32.002-31.995-48-80-48-112c0-48 37.594-74.993 80-128c22.35-27.937 32-144 0-208c48-16 48-16 96 0z"/>`),
		g.Group(children),
	)
}