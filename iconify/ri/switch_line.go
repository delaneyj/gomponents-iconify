package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwitchLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.001 3v18h-4.4a4.6 4.6 0 0 1-4.6-4.6V7.6a4.6 4.6 0 0 1 4.6-4.6h4.4Zm-2 2h-2.4a2.6 2.6 0 0 0-2.6 2.6v8.8a2.6 2.6 0 0 0 2.6 2.6h2.4V5Zm-2.5 5a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3Zm6.5-7h2.4a4.6 4.6 0 0 1 4.6 4.6v8.8a4.6 4.6 0 0 1-4.6 4.6h-2.4V3Zm3 11.7a1.8 1.8 0 1 0 0-3.6a1.8 1.8 0 0 0 0 3.6Z"/>`),
		g.Group(children),
	)
}