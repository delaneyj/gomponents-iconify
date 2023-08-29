package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonSimpleBikeLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M164 78a26 26 0 1 0-26-26a26 26 0 0 0 26 26Zm0-40a14 14 0 1 1-14 14a14 14 0 0 1 14-14Zm36 100a38 38 0 1 0 38 38a38 38 0 0 0-38-38Zm0 64a26 26 0 1 1 26-26a26 26 0 0 1-26 26ZM56 138a38 38 0 1 0 38 38a38 38 0 0 0-38-38Zm0 64a26 26 0 1 1 26-26a26 26 0 0 1-26 26Zm136-84h-40a6 6 0 0 1-4.24-1.76L120 88.49L96.49 112l35.75 35.76A6 6 0 0 1 134 152v48a6 6 0 0 1-12 0v-45.51l-38.24-38.25a6 6 0 0 1 0-8.48l32-32a6 6 0 0 1 8.48 0L154.49 106H192a6 6 0 0 1 0 12Z"/>`),
		g.Group(children),
	)
}