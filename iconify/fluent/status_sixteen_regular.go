package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StatusSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14.354 1.647a2.268 2.268 0 0 0-3.207 0l-4.25 4.25a.5.5 0 0 0-.121.195l-1.25 3.75a.5.5 0 0 0 .632.633l3.75-1.25a.5.5 0 0 0 .196-.121l4.25-4.25a2.268 2.268 0 0 0 0-3.207Zm-2.5.707a1.268 1.268 0 0 1 1.793 1.793L9.48 8.313l-2.69.897l.897-2.69l4.167-4.166ZM8 3c.123 0 .245.005.366.013l.883-.883a6 6 0 1 0 4.621 4.621l-.883.884A5 5 0 1 1 8 3Z"/>`),
		g.Group(children),
	)
}