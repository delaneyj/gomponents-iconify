package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StoolThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M196 64V40a12 12 0 0 0-12-12H72a12 12 0 0 0-12 12v24a12 12 0 0 0 12 12h11.32L60.05 223.38a4 4 0 0 0 3.33 4.62a4.46 4.46 0 0 0 .62 0a4 4 0 0 0 4-3.38L76.26 172h103.48l8.31 52.62A4 4 0 0 0 192 228a4.46 4.46 0 0 0 .63-.05a4 4 0 0 0 3.33-4.57L172.68 76H184a12 12 0 0 0 12-12ZM68 64V40a4 4 0 0 1 4-4h112a4 4 0 0 1 4 4v24a4 4 0 0 1-4 4H72a4 4 0 0 1-4-4Zm110.48 100h-101l13.9-88h73.16Z"/>`),
		g.Group(children),
	)
}