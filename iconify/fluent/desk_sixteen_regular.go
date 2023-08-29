package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeskSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2.75 3A1.75 1.75 0 0 0 1 4.75v6.5c0 .966.784 1.75 1.75 1.75h3.5A1.75 1.75 0 0 0 8 11.25V7h6v5.5a.5.5 0 0 0 1 0V4.75A1.75 1.75 0 0 0 13.25 3H2.75ZM2 7h5v4.25a.75.75 0 0 1-.75.75h-3.5a.75.75 0 0 1-.75-.75V7Zm0-1V4.75A.75.75 0 0 1 2.75 4h10.5a.75.75 0 0 1 .75.75V6H2Zm1.5 2a.5.5 0 0 0 0 1h2a.5.5 0 0 0 0-1h-2Z"/>`),
		g.Group(children),
	)
}