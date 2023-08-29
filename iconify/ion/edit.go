package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Edit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M163 439.573l-90.569-90.569L322.78 98.656l90.57 90.569z" fill="currentColor"/><path d="M471.723 88.393l-48.115-48.114c-11.723-11.724-31.558-10.896-44.304 1.85l-45.202 45.203 90.569 90.568 45.202-45.202c12.743-12.746 13.572-32.582 1.85-44.305z" fill="currentColor"/><path d="M64.021 363.252L32 480l116.737-32.021z" fill="currentColor"/>`),
		g.Group(children),
	)
}