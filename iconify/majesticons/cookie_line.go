package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CookieLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.419 12c0-6.4 5.358-8 8.037-8h1.005v4h3.014v3h4.019v1c0 6.4-5.024 8-8.038 8c-3.014 0-8.037-1.6-8.037-8zm5.023-4h.001m-2.01 4h0m5.023 1h.001m-2.01 3h0m6.028-1h0"/>`),
		g.Group(children),
	)
}