package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowOutlineDownLeftTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.666 17.99a1.5 1.5 0 0 1-1.657-1.657L3.088 6.63c.138-1.25 1.662-1.784 2.551-.895l1.067 1.067L11.07 2.44a1.5 1.5 0 0 1 2.122 0l4.37 4.37a1.5 1.5 0 0 1 0 2.121l-4.364 4.363l1.068 1.067c.889.89.354 2.413-.896 2.552L3.666 17.99Z"/>`),
		g.Group(children),
	)
}