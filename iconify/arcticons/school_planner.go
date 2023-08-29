package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SchoolPlanner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.272 5.525h-1.98c-.947 0-1.722.852-1.722 1.895v33.16c0 1.043.775 1.895 1.722 1.895h1.98m0-36.95v36.95h21.436c.947 0 1.721-.852 1.721-1.894V7.419c0-1.042-.775-1.894-1.721-1.894H14.272Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.549 5.846v7.206l3.69-2.13l3.723 2.149V5.613M20.185 28.299H33.01v-2.917H20.185v2.916Zm0 5.868h9.52v-2.836h-9.52v2.836Z"/>`),
		g.Group(children),
	)
}