package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sync(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m23.5 13.5l-3.086 3.086L19 18l-4.5-4.5l1.414-1.414L18 14.172V12c0-3.308-2.692-6-6-6V4a8 8 0 0 1 8 8v2.172l2.086-2.086L23.5 13.5zM6 12V9.828l2.086 2.086L9.5 10.5L5 6L3.586 7.414L.5 10.5l1.414 1.414L4 9.828V12a8 8 0 0 0 8 8v-2c-3.308 0-6-2.692-6-6z"/>`),
		g.Group(children),
	)
}