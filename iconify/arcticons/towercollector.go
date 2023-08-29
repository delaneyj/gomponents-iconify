package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Towercollector(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9 36.43h0a19.34 19.34 0 0 1 0-24.86m30 0a19.37 19.37 0 0 1 0 24.86m-24.74-4.32h0a12.57 12.57 0 0 1 0-16.14m19.56 0a12.57 12.57 0 0 1 0 16.14M24 17.84a6.16 6.16 0 0 1 2.3 11.87l-2.3 5l-2.3-5A6.16 6.16 0 0 1 24 17.84Zm0 3.69A2.47 2.47 0 1 0 26.47 24A2.44 2.44 0 0 0 24 21.53Z"/>`),
		g.Group(children),
	)
}