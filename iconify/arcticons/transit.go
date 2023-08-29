package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Transit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.446 7.63v6.961M39.554 40.37v-7.165m0-14.234v-3.549h0a7.777 7.777 0 1 0-15.554 0v17.156a7.777 7.777 0 1 1-15.554 0v-3.753"/><ellipse cx="8.446" cy="21.708" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3.946" ry="3.955"/><ellipse cx="39.554" cy="26.088" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3.946" ry="3.955"/>`),
		g.Group(children),
	)
}