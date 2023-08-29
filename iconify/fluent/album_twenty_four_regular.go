package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlbumTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 9a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v1a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2V9Zm2-.5a.5.5 0 0 0-.5.5v1a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5V9a.5.5 0 0 0-.5-.5h-4ZM4 4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2H4Zm16 1.5a.5.5 0 0 1 .5.5v12a.5.5 0 0 1-.5.5H7.5v-13H20Zm-16 0h2v13H4a.5.5 0 0 1-.5-.5V6a.5.5 0 0 1 .5-.5Z"/>`),
		g.Group(children),
	)
}