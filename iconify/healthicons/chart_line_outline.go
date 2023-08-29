package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartLineOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6 41a1 1 0 0 0 1 1h34v-2H8v-4.641l6.352-7.713a3.998 3.998 0 0 0 4.976-1.426l4.675 1.948a4 4 0 1 0 7.383-2.3l4.975-6.218A4 4 0 1 0 34.8 18.4l-4.875 6.093a4 4 0 0 0-5.489 1.689l-4.45-1.854a4 4 0 1 0-7.193 2.064L8 32.212V7H6v34Zm12-17a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm10 6a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm12-14a2 2 0 1 1-4 0a2 2 0 0 1 4 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}