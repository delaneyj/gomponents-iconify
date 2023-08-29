package formkit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Avatarman(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M4 10h5c1.66 0 3 1.34 3 3v2H1v-2c0-1.66 1.34-3 3-3Zm0-6h5v2.5a2.5 2.5 0 0 1-5 0V4Zm5 0c.55 0 1 .45 1 1s-.45 1-1 1V4ZM4 6c-.55 0-1-.45-1-1s.45-1 1-1v2Z"/><path fill="currentColor" d="M4.12 4.12h-.5C2.87 3.5 3 2.55 3 1.75s.5-.26 1-.26s1 .5 1 .5c-.88 0-1-1-1-1h3c1.1 0 2 .9 2 2h.5s.25.75-.12 1.25l-5.25-.12Z"/>`),
		g.Group(children),
	)
}