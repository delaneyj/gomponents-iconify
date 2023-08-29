package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PushPinSimpleThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 172h-12.64L180.77 44H192a4 4 0 0 0 0-8H64a4 4 0 0 0 0 8h11.23L52.64 172H40a4 4 0 0 0 0 8h84v60a4 4 0 0 0 8 0v-60h84a4 4 0 0 0 0-8ZM83.36 44h89.28l22.59 128H60.77Z"/>`),
		g.Group(children),
	)
}