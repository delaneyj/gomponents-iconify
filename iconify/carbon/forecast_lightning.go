package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ForecastLightning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m15.67 24l-1.736-1l2.287-4h-3.889l3.993-7l1.737 1l-2.284 4h3.89l-3.998 7z"/><path fill="currentColor" d="M26 18A10 10 0 1 1 16 8h4v5l6-6l-6-6v5h-4a12 12 0 1 0 12 12Z"/>`),
		g.Group(children),
	)
}