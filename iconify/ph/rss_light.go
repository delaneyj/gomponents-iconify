package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RssLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M97.5 158.5A69.54 69.54 0 0 1 118 208a6 6 0 0 1-12 0a58 58 0 0 0-58-58a6 6 0 0 1 0-12a69.54 69.54 0 0 1 49.5 20.5ZM48 90a6 6 0 0 0 0 12a106 106 0 0 1 106 106a6 6 0 0 0 12 0A118 118 0 0 0 48 90Zm117.38.62A164.92 164.92 0 0 0 48 42a6 6 0 0 0 0 12a153 153 0 0 1 108.89 45.11A153 153 0 0 1 202 208a6 6 0 0 0 12 0a164.92 164.92 0 0 0-48.62-117.38ZM52 194a10 10 0 1 0 10 10a10 10 0 0 0-10-10Z"/>`),
		g.Group(children),
	)
}