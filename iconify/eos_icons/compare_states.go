package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CompareStates(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 16.184V6a2 2 0 0 0-2-2h-5.172l1.586-1.586L14 1l-2.586 2.586L10 5l1.414 1.414L14 9l1.414-1.414L13.828 6H19v10.184a3 3 0 1 0 2 0Zm-8.414 1.402L10 15l-1.414 1.414L10.172 18H5V7.816a3 3 0 1 0-2 0V18a2 2 0 0 0 2 2h5.172l-1.586 1.586L10 23l2.586-2.586L14 19l-1.414-1.414Z"/>`),
		g.Group(children),
	)
}