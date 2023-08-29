package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlbumAddTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 6a2 2 0 0 1 2-2h1v7.174a6.488 6.488 0 0 0-3 1.636V6Zm14 2.5h-4a.5.5 0 0 0-.5.5v1a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5V9a.5.5 0 0 0-.5-.5ZM12.502 20H20a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2H6.5v7a6.5 6.5 0 0 1 6.002 9ZM10 9a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v1a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2V9Zm-3.5 3a5.5 5.5 0 1 1 0 11a5.5 5.5 0 0 1 0-11Zm.501 8.503V18h2.496a.5.5 0 0 0 0-1H7v-2.5a.5.5 0 1 0-1 0V17H3.496a.5.5 0 0 0 0 1H6v2.503a.5.5 0 1 0 1 0Z"/>`),
		g.Group(children),
	)
}