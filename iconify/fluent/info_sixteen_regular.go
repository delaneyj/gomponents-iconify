package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InfoSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 7a.5.5 0 0 1 .5.5v3a.5.5 0 0 1-1 0v-3A.5.5 0 0 1 8 7Zm0-.75a.749.749 0 1 0 0-1.5a.749.749 0 0 0 0 1.498ZM2 8a6 6 0 1 1 12 0A6 6 0 0 1 2 8Zm6-5a5 5 0 1 0 0 10A5 5 0 0 0 8 3Z"/>`),
		g.Group(children),
	)
}