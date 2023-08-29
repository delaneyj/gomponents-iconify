package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SankeyDiagramAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 2H2v14h6.111a7.036 7.036 0 0 1 3.13.74L15.764 19l-4.522 2.26a7.036 7.036 0 0 1-3.13.74H2v8h7.223a9.045 9.045 0 0 0 4.025-.95l8.622-4.31A7.036 7.036 0 0 1 25 24h5V14h-5a7.036 7.036 0 0 1-3.13-.74L15.348 10H30ZM8 4h16v4H8ZM4 4h2v10H4Zm0 20h2v4H4Zm16.975-1.05l-8.622 4.31a7.036 7.036 0 0 1-3.13.74H8v-4h.111a9.045 9.045 0 0 0 4.025-.95L18 20.118l1.864.932a9.045 9.045 0 0 0 4.025.95H24v.058a9.052 9.052 0 0 0-3.025.892ZM28 22h-2v-6h2ZM12.354 10.74l8.621 4.31a9.052 9.052 0 0 0 3.025.89V20h-.111a7.036 7.036 0 0 1-3.13-.74l-8.622-4.31A9.045 9.045 0 0 0 8.11 14H8v-4h1.223a7.036 7.036 0 0 1 3.13.74ZM28 8h-2V4h2Z"/>`),
		g.Group(children),
	)
}