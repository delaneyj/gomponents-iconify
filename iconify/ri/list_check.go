package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 4h13v2H8V4Zm-5-.5h3v3H3v-3Zm0 7h3v3H3v-3Zm0 7h3v3H3v-3ZM8 11h13v2H8v-2Zm0 7h13v2H8v-2Z"/>`),
		g.Group(children),
	)
}