package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentBulletListSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5 1a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2V5.414a1.5 1.5 0 0 0-.44-1.06L9.647 1.439A1.5 1.5 0 0 0 8.586 1H5ZM4 3a1 1 0 0 1 1-1h3v2.5A1.5 1.5 0 0 0 9.5 6H12v7a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V3Zm7.793 2H9.5a.5.5 0 0 1-.5-.5V2.207L11.793 5ZM7 7.5a.5.5 0 0 1 .5-.5h3a.5.5 0 0 1 0 1h-3a.5.5 0 0 1-.5-.5ZM7.5 9a.5.5 0 0 0 0 1h3a.5.5 0 0 0 0-1h-3ZM7 11.5a.5.5 0 0 1 .5-.5h3a.5.5 0 0 1 0 1h-3a.5.5 0 0 1-.5-.5ZM5.5 8a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1ZM6 9.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0ZM5.5 12a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Z"/>`),
		g.Group(children),
	)
}