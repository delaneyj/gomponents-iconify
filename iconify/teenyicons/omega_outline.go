package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OmegaOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M9.333 11.687a.5.5 0 1 0 .334.943l-.334-.943Zm-4 .943a.5.5 0 1 0 .334-.943l-.334.943ZM5.5 14.5v.5H6v-.5h-.5Zm4 0H9v.5h.5v-.5ZM7.5 1A5.5 5.5 0 0 1 13 6.5h1A6.5 6.5 0 0 0 7.5 0v1Zm0-1A6.5 6.5 0 0 0 1 6.5h1A5.5 5.5 0 0 1 7.5 1V0ZM13 6.5a5.503 5.503 0 0 1-3.667 5.187l.334.943A6.503 6.503 0 0 0 14 6.5h-1Zm-7.333 5.187A5.503 5.503 0 0 1 2 6.5H1c0 2.83 1.81 5.238 4.333 6.13l.334-.943ZM0 15h5.5v-1H0v1Zm6-.5V12H5v2.5h1Zm9-.5H9.5v1H15v-1Zm-5 .5V12H9v2.5h1Z"/>`),
		g.Group(children),
	)
}