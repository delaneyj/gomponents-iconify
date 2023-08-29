package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SearchFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M20.5 6C12.492 6 6 12.492 6 20.5S12.492 35 20.5 35a14.44 14.44 0 0 0 9.138-3.241l9.801 9.801a1.5 1.5 0 1 0 2.121-2.121l-9.8-9.801A14.44 14.44 0 0 0 35 20.5C35 12.492 28.508 6 20.5 6ZM9 20.5C9 14.149 14.149 9 20.5 9S32 14.149 32 20.5S26.851 32 20.5 32S9 26.851 9 20.5Z"/>`),
		g.Group(children),
	)
}