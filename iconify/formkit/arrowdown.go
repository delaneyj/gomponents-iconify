package formkit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Arrowdown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M4.5 13c-.28 0-.5-.22-.5-.5v-9c0-.28.22-.5.5-.5s.5.22.5.5v9c0 .28-.22.5-.5.5Z"/><path fill="currentColor" d="M4.5 14a.47.47 0 0 1-.35-.15l-3.5-3.5c-.2-.2-.2-.51 0-.71c.2-.2.51-.2.71 0l3.15 3.15l3.15-3.15c.2-.2.51-.2.71 0c.2.2.2.51 0 .71l-3.5 3.5c-.1.1-.23.15-.35.15Z"/>`),
		g.Group(children),
	)
}