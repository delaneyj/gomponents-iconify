package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clickup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m2 18.439l3.69-2.828c1.961 2.56 4.044 3.739 6.363 3.739c2.307 0 4.33-1.166 6.203-3.704L22 18.405C19.298 22.065 15.941 24 12.053 24C8.178 24 4.788 22.078 2 18.439zM12.04 6.15l-6.568 5.66l-3.036-3.52L12.055 0l9.543 8.296l-3.05 3.509z"/>`),
		g.Group(children),
	)
}