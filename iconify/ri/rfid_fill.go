package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RfidFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.364 18.364a9 9 0 0 0 0-12.728l1.414-1.414c4.296 4.295 4.296 11.26 0 15.556l-1.414-1.414ZM5.636 5.636a9 9 0 0 0 0 12.728l-1.414 1.414c-4.296-4.296-4.296-11.26 0-15.556l1.414 1.414Zm9.9 9.9a5 5 0 0 0 0-7.072L16.95 7.05a7 7 0 0 1 0 9.9l-1.414-1.415ZM8.463 8.463a5 5 0 0 0 0 7.071L7.05 16.95a7 7 0 0 1 0-9.9l1.414 1.414ZM12 14a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/>`),
		g.Group(children),
	)
}