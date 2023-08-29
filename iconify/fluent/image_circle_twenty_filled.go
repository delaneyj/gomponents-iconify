package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageCircleTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18 10a7.97 7.97 0 0 1-1.998 5.29l-4.95-4.871a1.5 1.5 0 0 0-2.104 0l-4.95 4.871A8 8 0 1 1 18 10Zm-7.649 1.131l4.944 4.866A7.97 7.97 0 0 1 10 18a7.97 7.97 0 0 1-5.295-2.003l4.945-4.866a.5.5 0 0 1 .701 0ZM12.75 8.5a1.25 1.25 0 1 0 0-2.5a1.25 1.25 0 0 0 0 2.5Z"/>`),
		g.Group(children),
	)
}