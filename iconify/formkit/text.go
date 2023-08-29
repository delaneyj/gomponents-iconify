package formkit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Text(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M9.5 4h-8c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h8c.28 0 .5.22.5.5s-.22.5-.5.5Zm4 3h-6c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h6c.28 0 .5.22.5.5s-.22.5-.5.5Z"/><path fill="currentColor" d="M5.5 14c-.28 0-.5-.22-.5-.5V3.55c0-.28.22-.5.5-.5s.5.22.5.5v9.95c0 .28-.22.5-.5.5Zm5 0c-.28 0-.5-.22-.5-.5V6.55c0-.28.22-.5.5-.5s.5.22.5.5v6.95c0 .28-.22.5-.5.5Z"/>`),
		g.Group(children),
	)
}