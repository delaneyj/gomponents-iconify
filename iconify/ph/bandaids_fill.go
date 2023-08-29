package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BandaidsFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><circle cx="128" cy="128" r="12"/><path d="m184.57 128l27.71-27.72a40 40 0 1 0-56.56-56.56L128 71.43l-27.72-27.71a40 40 0 1 0-56.56 56.56L71.43 128l-27.71 27.72a40 40 0 1 0 56.56 56.56L128 184.57l27.72 27.71a40 40 0 1 0 56.56-56.56ZM128 94.06L161.94 128L128 161.94L94.06 128Zm39-39A24 24 0 1 1 201 89l-27.72 27.72l-33.97-33.97ZM89 201a24 24 0 1 1-34-34l27.72-27.72l33.94 33.94Z"/></g>`),
		g.Group(children),
	)
}