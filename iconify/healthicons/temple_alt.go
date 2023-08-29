package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TempleAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M36 11.048a1 1 0 0 0-1-1h-5.083a6.002 6.002 0 0 0-11.834 0H13a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h1v2h-2a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1v2h-2a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h1V36H8a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h12v.048h8V42h12a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1h-3v-3.952h1a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1h-2v-2a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1h-2v-2h1a1 1 0 0 0 1-1v-4Zm-22 1h7v2h-7v-2Zm7 4h-5v2h5v-2Zm2 2v-2h2v2h-2Zm-2 2h-8v2h8v-2Zm2 2v-2h2v2h-2Zm-2 2h-7v2h7v-2Zm2 2v-2h2v2h-2Zm-12 2v2h10v-2H11Zm26 2v-2H27v2h10Zm-12-2V30h-2v-1.952h2Zm2-2h7v-2h-7v2Zm0-4h8v-2h-8v2Zm0-4h5v-2h-5v2Zm0-4h7v-2h-7v2Zm-2-2v2h-2v-2h2ZM35 36v-3.952H13V36h7v-1.952h8V36h7Zm-7 2v2h11v-2H28Zm-8 0v2H9v-2h11Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}