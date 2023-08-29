package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HardDriveTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.44 4a2.25 2.25 0 0 0-2.025 1.27L2.73 10.821c.465-.206.98-.321 1.521-.321h15.5a3.74 3.74 0 0 1 1.52.321L18.586 5.27A2.25 2.25 0 0 0 16.559 4H7.441Zm12.31 8A2.25 2.25 0 0 1 22 14.25v2.5A2.25 2.25 0 0 1 19.75 19H4.25A2.25 2.25 0 0 1 2 16.75v-2.5A2.25 2.25 0 0 1 4.25 12h15.5Zm-1.25 4.5a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/>`),
		g.Group(children),
	)
}