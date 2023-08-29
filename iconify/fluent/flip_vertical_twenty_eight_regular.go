package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlipVerticalTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M23.654 2.118A.75.75 0 0 1 24 2.75v9.5a.75.75 0 0 1-.75.75H2.75a.75.75 0 0 1-.315-1.43l20.5-9.5a.75.75 0 0 1 .719.048ZM6.152 11.5H22.5V3.924L6.152 11.5ZM24 25.25a.75.75 0 0 1-1.065.68l-20.5-9.5A.75.75 0 0 1 2.75 15h20.5a.75.75 0 0 1 .75.75v9.5Z"/>`),
		g.Group(children),
	)
}