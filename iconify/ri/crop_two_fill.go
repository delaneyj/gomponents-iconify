package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CropTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17.586 5l2.556-2.556l1.414 1.414L19 6.414V17h3v2h-3v3h-2V7H9V5h8.586ZM15 17v2H6a1 1 0 0 1-1-1V7H2V5h3V2h2v15h8ZM9 9h6v6H9V9Z"/>`),
		g.Group(children),
	)
}