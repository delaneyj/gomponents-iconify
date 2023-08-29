package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RulerF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16.757 9.243l-1.414 1.414l-2.828-2.829A1 1 0 1 0 11.1 9.243l2.828 2.828l-1.414 1.414l-2.122-2.121a1 1 0 1 0-1.414 1.414l2.122 2.121l-1.415 1.415l-2.828-2.829A1 1 0 0 0 5.444 14.9l2.828 2.829l-1.414 1.414a2 2 0 0 1-2.829 0l-2.828-2.828a2 2 0 0 1 0-2.829L13.929.757a2 2 0 0 1 2.828 0l2.829 2.829a2 2 0 0 1 0 2.828l-1.414 1.414l-2.122-2.12a1 1 0 1 0-1.414 1.413l2.121 2.122z"/>`),
		g.Group(children),
	)
}