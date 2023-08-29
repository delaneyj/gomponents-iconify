package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PentagonFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M21.502 4.313a4.25 4.25 0 0 1 5 0l15.75 11.456a4.25 4.25 0 0 1 1.543 4.75l-6.018 18.54a4.25 4.25 0 0 1-4.042 2.938H14.268a4.25 4.25 0 0 1-4.042-2.938l-6.018-18.54a4.25 4.25 0 0 1 1.542-4.75L21.502 4.313Z"/>`),
		g.Group(children),
	)
}