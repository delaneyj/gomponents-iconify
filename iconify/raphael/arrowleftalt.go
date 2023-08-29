package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Arrowleftalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 30.534c8.027 0 14.534-6.507 14.534-14.534c0-8.027-6.507-14.534-14.534-14.534C7.973 1.466 1.466 7.973 1.466 16c0 8.027 6.507 14.534 14.534 14.534zm2.335-24.258l3.536 3.538L15.685 16l6.187 6.188l-3.535 3.537L8.612 16l9.723-9.724z"/>`),
		g.Group(children),
	)
}