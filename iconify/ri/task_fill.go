package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TaskFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 3v18.008a.993.993 0 0 1-.993.992H3.993A1 1 0 0 1 3 21.008V2.992C3 2.444 3.445 2 3.993 2H20a1 1 0 0 1 1 1Zm-9.707 10.121l-2.475-2.475l-1.414 1.415l3.889 3.889l5.657-5.657l-1.414-1.414l-4.243 4.242Z"/>`),
		g.Group(children),
	)
}