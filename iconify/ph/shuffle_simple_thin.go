package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShuffleSimpleThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M212 48v40a4 4 0 0 1-8 0V57.66l-50.54 50.54a4 4 0 1 1-5.66-5.66L198.34 52H168a4 4 0 0 1 0-8h40a4 4 0 0 1 4 4Zm-4 116a4 4 0 0 0-4 4v30.34L50.83 45.17a4 4 0 0 0-5.66 5.66L198.34 204H168a4 4 0 0 0 0 8h40a4 4 0 0 0 4-4v-40a4 4 0 0 0-4-4Zm-105.46-16.2l-57.37 57.37a4 4 0 0 0 5.66 5.66l57.37-57.37a4 4 0 1 0-5.66-5.66Z"/>`),
		g.Group(children),
	)
}