package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FindInPage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M334.627 16H48v480h424V153.373ZM440 464H80V48h241.373L440 166.627Z"/><path fill="currentColor" d="M239.861 152a95.861 95.861 0 1 0 53.624 175.284l68.03 68.029l22.627-22.626l-67.5-67.5A95.816 95.816 0 0 0 239.861 152ZM176 247.861a63.862 63.862 0 1 1 63.861 63.861A63.933 63.933 0 0 1 176 247.861Z"/>`),
		g.Group(children),
	)
}