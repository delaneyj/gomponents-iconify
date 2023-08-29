package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ruby(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.156.083c3.033.525 3.893 2.598 3.829 4.77L24 4.822L22.635 22.71L4.89 23.926h.016C3.433 23.864.15 23.729 0 19.139l1.645-3l2.819 6.586l.503 1.172l2.805-9.144l-.03.007l.016-.03l9.255 2.956l-1.396-5.431l-.99-3.9l8.82-.569l-.615-.51L16.5 2.114L20.159.073l-.003.01zM0 19.089zM5.13 5.073c3.561-3.533 8.157-5.621 9.922-3.84c1.762 1.777-.105 6.105-3.673 9.636c-3.563 3.532-8.103 5.734-9.864 3.957c-1.766-1.777.045-6.217 3.612-9.75l.003-.003z"/>`),
		g.Group(children),
	)
}