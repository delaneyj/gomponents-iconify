package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IceCreamThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M204 92.4V88a76 76 0 0 0-152 0v4.4a20 20 0 0 0 4 39.6h5.68l62.85 110a4 4 0 0 0 6.94 0l62.85-110H200a20 20 0 0 0 4-39.6Zm-76 139.54L70.89 132h28.5l42.86 75Zm9.11-99.94l24 42l-14.25 24.94L108.61 132Zm28.6 33.94L146.32 132h38.79ZM200 124H56a12 12 0 0 1 0-24a4 4 0 0 0 4-4v-8a68 68 0 0 1 136 0v8a4 4 0 0 0 4 4a12 12 0 0 1 0 24Z"/>`),
		g.Group(children),
	)
}