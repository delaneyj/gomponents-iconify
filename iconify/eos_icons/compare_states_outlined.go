package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CompareStatesOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10 15l-1.414 1.414L10.172 18H5V7.816a3 3 0 1 0-2 0V18a2 2 0 0 0 2 2h5.172l-1.586 1.586L10 23l4-4ZM4 6a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm17 10.184V6a2 2 0 0 0-2-2h-5.172l1.586-1.586L14 1l-4 4l4 4l1.414-1.414L13.828 6H19v10.184a3 3 0 1 0 2 0ZM20 20a1 1 0 1 1 1-1a1 1 0 0 1-1 1Z"/>`),
		g.Group(children),
	)
}