package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9.38 4.677a1.25 1.25 0 0 0-1.88 1.08v16.488a1.25 1.25 0 0 0 1.88 1.079l14.698-8.59a.85.85 0 0 0 0-1.467L9.381 4.677ZM6 5.757c0-2.124 2.304-3.447 4.138-2.375l14.697 8.59c1.552.907 1.552 3.15 0 4.057l-14.697 8.59C8.304 25.691 6 24.369 6 22.245V5.756Z"/>`),
		g.Group(children),
	)
}