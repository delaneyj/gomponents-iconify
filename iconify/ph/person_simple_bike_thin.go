package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonSimpleBikeThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M164 76a24 24 0 1 0-24-24a24 24 0 0 0 24 24Zm0-40a16 16 0 1 1-16 16a16 16 0 0 1 16-16Zm36 104a36 36 0 1 0 36 36a36 36 0 0 0-36-36Zm0 64a28 28 0 1 1 28-28a28 28 0 0 1-28 28ZM56 140a36 36 0 1 0 36 36a36 36 0 0 0-36-36Zm0 64a28 28 0 1 1 28-28a28 28 0 0 1-28 28Zm136-88h-40a4 4 0 0 1-2.83-1.17L120 85.66L93.66 112l37.17 37.17A4 4 0 0 1 132 152v48a4 4 0 0 1-8 0v-46.34l-38.83-38.83a4 4 0 0 1 0-5.66l32-32a4 4 0 0 1 5.66 0L153.66 108H192a4 4 0 0 1 0 8Z"/>`),
		g.Group(children),
	)
}