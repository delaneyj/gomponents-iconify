package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AssetConfirm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m21 27.18l-2.59-2.59L17 26l4 4l7-7l-1.41-1.41L21 27.18zM12 24a4 4 0 1 1 4-4a4.004 4.004 0 0 1-4 4zm0-6a2 2 0 1 0 2 2a2.002 2.002 0 0 0-2-2z"/><path fill="currentColor" d="M26 2a3.86 3.86 0 0 0-1.85.48L7.76 10.96A9.99 9.99 0 0 0 12 30a9.345 9.345 0 0 0 2-.21v-2.04a8.229 8.229 0 0 1-2 .25a8 8 0 1 1 8-8a8.266 8.266 0 0 1-.06 1h2.78l6.84-13.21A3.973 3.973 0 0 0 26 2Zm-4.14 16.34a10.019 10.019 0 0 0-8.22-8.19l8.38-4.34c0 .06-.02.12-.02.19a3.999 3.999 0 0 0 4 4c.06 0 .12-.02.19-.02ZM26 8a2 2 0 1 1 2-2a2.006 2.006 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}