package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlipVerticalThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M27.43 2.201c.356.23.57.625.57 1.049v10.5c0 .69-.56 1.25-1.25 1.25H3.25a1.25 1.25 0 0 1-.51-2.391l23.5-10.5a1.25 1.25 0 0 1 1.19.092ZM9.112 12.5H25.5V5.178L9.112 12.5ZM28 29a1 1 0 0 1-1.416.91l-24-11A1 1 0 0 1 3 17h24a1 1 0 0 1 1 1v11Z"/>`),
		g.Group(children),
	)
}