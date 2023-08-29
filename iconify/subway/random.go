package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Random(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M341.3 28.3v85.3H128c-70.7 0-128 57.3-128 128c0 21.5 5.8 41.4 15.2 59.2L68 263.2c-2.4-6.8-4-13.9-4-21.5c0-35.4 28.7-64 64-64h213.3V263L512 156.3V135L341.3 28.3zM444 262.8c2.4 6.8 4 13.9 4 21.5c0 35.4-28.6 64-64 64H170.7V263L0 369.7V391l170.7 106.7v-85.3H384c70.7 0 128-57.3 128-128c0-21.5-5.8-41.4-15.2-59.2L444 262.8z"/>`),
		g.Group(children),
	)
}