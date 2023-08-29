package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EscalatorUpThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 44h-52.5a12 12 0 0 0-8.81 3.86l-91.25 98.85A4 4 0 0 1 68.5 148H32a12 12 0 0 0-12 12v40a12 12 0 0 0 12 12h52.5a12 12 0 0 0 8.81-3.86l91.25-98.85a4 4 0 0 1 2.94-1.29H224a12 12 0 0 0 12-12V56a12 12 0 0 0-12-12Zm4 52a4 4 0 0 1-4 4h-36.5a12 12 0 0 0-8.81 3.86l-91.25 98.85A4 4 0 0 1 84.5 204H32a4 4 0 0 1-4-4v-40a4 4 0 0 1 4-4h36.5a12 12 0 0 0 8.81-3.86l91.25-98.85A4 4 0 0 1 171.5 52H224a4 4 0 0 1 4 4Zm-1.17 77.17a4 4 0 0 1-5.66 5.66L204 161.66V208a4 4 0 0 1-8 0v-46.34l-17.17 17.17a4 4 0 0 1-5.66-5.66l24-24a4 4 0 0 1 5.66 0Z"/>`),
		g.Group(children),
	)
}