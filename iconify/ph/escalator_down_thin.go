package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EscalatorDownThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M173.17 82.83a4 4 0 0 1 5.66-5.66L196 94.34V48a4 4 0 0 1 8 0v46.34l17.17-17.17a4 4 0 1 1 5.66 5.66l-24 24a4 4 0 0 1-5.66 0ZM236 160v40a12 12 0 0 1-12 12h-52.5a12.06 12.06 0 0 1-8.82-3.86l-91.24-98.85A4 4 0 0 0 68.5 108H32a12 12 0 0 1-12-12V56a12 12 0 0 1 12-12h52.5a12.06 12.06 0 0 1 8.82 3.86l91.24 98.85a4 4 0 0 0 2.94 1.29H224a12 12 0 0 1 12 12Zm-8 0a4 4 0 0 0-4-4h-36.5a12.06 12.06 0 0 1-8.82-3.86L87.44 53.29A4 4 0 0 0 84.5 52H32a4 4 0 0 0-4 4v40a4 4 0 0 0 4 4h36.5a12.06 12.06 0 0 1 8.82 3.86l91.24 98.85a4 4 0 0 0 2.94 1.29H224a4 4 0 0 0 4-4Z"/>`),
		g.Group(children),
	)
}