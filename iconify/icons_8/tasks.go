package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tasks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 5v6h22V5H5zm6 2h14v2H11V7zm-6 6v6h22v-6H5zm16 2h4v2h-4v-2zM5 21v6h22v-6H5zm11 2h9v2h-9v-2z"/>`),
		g.Group(children),
	)
}