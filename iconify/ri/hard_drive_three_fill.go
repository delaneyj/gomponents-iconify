package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HardDriveThreeFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.5 2a1 1 0 0 0-.992.876l-1.5 12A1 1 0 0 0 3 15v6a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1v-6c0-.041-.003-.083-.008-.124l-1.5-12A1 1 0 0 0 18.5 2h-13ZM5 16h14v4H5v-4Zm10 1h2v2h-2v-2Zm-2 0h-2v2h2v-2Z"/>`),
		g.Group(children),
	)
}