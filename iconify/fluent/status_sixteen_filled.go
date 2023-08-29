package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StatusSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14.354 1.647a2.268 2.268 0 0 0-3.207 0l-4.25 4.25a.5.5 0 0 0-.121.195l-1.25 3.75a.5.5 0 0 0 .632.633l3.75-1.25a.5.5 0 0 0 .196-.121l4.25-4.25a2.268 2.268 0 0 0 0-3.207Zm-1.367 5.988a5 5 0 1 1-4.621-4.621l.883-.884a6 6 0 1 0 4.621 4.621l-.883.884Z"/>`),
		g.Group(children),
	)
}