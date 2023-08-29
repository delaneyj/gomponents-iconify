package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AssetView(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="22" cy="24" r="2" fill="currentColor"/><path fill="currentColor" d="M29.777 23.479A8.64 8.64 0 0 0 22 18a8.64 8.64 0 0 0-7.777 5.479L14 24l.223.521A8.64 8.64 0 0 0 22 30a8.64 8.64 0 0 0 7.777-5.479L30 24zM22 28a4 4 0 1 1 4-4a4.005 4.005 0 0 1-4 4zm-10-4a4 4 0 1 1 2.981-6.667l-1.49 1.334A2 2 0 1 0 12 22z"/><path fill="currentColor" d="M26 2a3.86 3.86 0 0 0-1.85.48L7.76 10.96A9.99 9.99 0 0 0 12 30v-2a8 8 0 1 1 5.63-13.68l1.4-1.4a10.006 10.006 0 0 0-5.39-2.77l8.38-4.34c0 .06-.02.12-.02.19a3.999 3.999 0 0 0 4 4c.06 0 .12-.02.19-.02L23.07 16h2.24l4.25-8.21A3.973 3.973 0 0 0 26 2Zm0 6a2 2 0 1 1 2-2a2.006 2.006 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}