package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InsertTable(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 5v22h22V5H5zm2 2h5v5H7V7zm7 0h4v5h-4V7zm6 0h5v5h-5V7zM7 14h5v4H7v-4zm7 0h4v4h-4v-4zm6 0h5v4h-5v-4zM7 20h5v5H7v-5zm7 0h4v5h-4v-5zm6 0h5v5h-5v-5z"/>`),
		g.Group(children),
	)
}