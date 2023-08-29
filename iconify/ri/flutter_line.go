package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlutterLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.598 10.684h2.828l-5.657 5.658l5.657 5.656h-2.828L8.94 16.341l5.657-5.657Zm-.194-8.68h2.829L5.919 13.318l-1.414-1.414l9.9-9.9Z"/>`),
		g.Group(children),
	)
}