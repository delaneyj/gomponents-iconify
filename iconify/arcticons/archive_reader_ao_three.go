package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArchiveReaderAoThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 5.5v37m17.7-29.1H6.3M34 33.8H14m10-20.4c-4.4 8-11.6 16.3-17.3 21.1M24 13.4c4.4 8 11.6 16.3 17.3 21.1"/>`),
		g.Group(children),
	)
}