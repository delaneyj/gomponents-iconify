package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stroller(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.5 3a8.5 8.5 0 0 0-7.212 13m14.424 0A8.46 8.46 0 0 0 20 11.5v-2h2.5M8 21a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm7 0a2 2 0 1 1 0-4a2 2 0 0 1 0 4ZM11.5 3v9m-8 0h16"/>`),
		g.Group(children),
	)
}