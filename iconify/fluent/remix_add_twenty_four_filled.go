package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RemixAddTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 3a1 1 0 0 1 1-1h9c5.523 0 10 4.477 10 10a9.968 9.968 0 0 1-2.859 7h-3.265A8.003 8.003 0 0 0 12 4H3a1 1 0 0 1-1-1Zm8.182 18.835c.59.108 1.197.165 1.818.165h9a1 1 0 1 0 0-2h-9a7.99 7.99 0 0 1-6.54-3.391A7.963 7.963 0 0 1 4 12a7.998 7.998 0 0 1 4.124-7H4.859A9.968 9.968 0 0 0 2 12c0 4.902 3.527 8.98 8.182 9.835ZM12 8a1 1 0 0 1 1 1v2h2a1 1 0 1 1 0 2h-2v2a1 1 0 1 1-2 0v-2H9a1 1 0 1 1 0-2h2V9a1 1 0 0 1 1-1Z"/>`),
		g.Group(children),
	)
}