package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Reblog(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22.086 9.914L20 7.828V18a2 2 0 0 1-2 2h-7v-2h7V7.828l-2.086 2.086L14.5 8.5L19 4l4.5 4.5l-1.414 1.414zM6 16.172V6h7V4H6a2 2 0 0 0-2 2v10.172l-2.086-2.086L.5 15.5L5 20l4.5-4.5l-1.414-1.414L6 16.172z"/>`),
		g.Group(children),
	)
}