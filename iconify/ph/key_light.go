package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M160 18a78.05 78.05 0 0 0-73.8 103.31l-58.44 58.45A6 6 0 0 0 26 184v40a6 6 0 0 0 6 6h40a6 6 0 0 0 6-6v-18h18a6 6 0 0 0 6-6v-18h18a6 6 0 0 0 4.24-1.76l10.45-10.44A78 78 0 1 0 160 18Zm0 144a65.63 65.63 0 0 1-24.43-4.67a6 6 0 0 0-6.64 1.26L117.51 170H96a6 6 0 0 0-6 6v18H72a6 6 0 0 0-6 6v18H38v-31.51l59.41-59.42a6 6 0 0 0 1.26-6.64A66 66 0 1 1 160 162Zm30-86a10 10 0 1 1-10-10a10 10 0 0 1 10 10Z"/>`),
		g.Group(children),
	)
}