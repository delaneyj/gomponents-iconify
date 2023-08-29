package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CropLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 17v2H6a1 1 0 0 1-1-1V7H2V5h3V2h2v15h8Zm2 5V7H9V5h9a1 1 0 0 1 1 1v11h3v2h-3v3h-2Z"/>`),
		g.Group(children),
	)
}