package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M23 23v466h466v-18H41v-82.184l85.854-57.234l70.023 70.022l65.133-260.536L387.28 203.7l67.79-107.97l19.317 11.858l6.102-71.1l-60.644 37.616l19.884 12.207l-59.01 93.99l-130.732-65.366l-62.865 251.462l-57.98-57.978L41 367.184V23H23z"/>`),
		g.Group(children),
	)
}