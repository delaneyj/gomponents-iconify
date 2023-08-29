package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExclamationSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0zM381.958 118.213h436.084l-93.75 660.571H475.708l-93.75-660.571zm220.386 705.542c71.262 0 129.053 57.79 129.053 129.053s-57.791 128.979-129.053 128.979s-129.053-57.717-129.053-128.979s57.791-129.053 129.053-129.053z"/>`),
		g.Group(children),
	)
}