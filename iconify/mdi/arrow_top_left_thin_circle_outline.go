package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTopLeftThinCircleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3.97c4.41 0 8.03 3.62 8.03 8.03c0 4.41-3.62 8.03-8.03 8.03c-4.41 0-8.03-3.62-8.03-8.03c0-4.41 3.62-8.03 8.03-8.03M12 2C6.46 2 2 6.46 2 12s4.46 10 10 10s10-4.46 10-10S17.54 2 12 2m-1.88 9.53L8 13.64V8h5.64l-2.11 2.12l4.97 4.98l-1.4 1.4"/>`),
		g.Group(children),
	)
}