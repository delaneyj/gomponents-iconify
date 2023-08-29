package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneInSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.566 19.152A18.466 18.466 0 0 1 5 9.86v-.002l-.519-1.14a3.5 3.5 0 0 1 1.237-4.355l1.333-.894a1 1 0 0 1 1.335.203l2.43 3.012a1 1 0 0 1-.183 1.431l-1.257.932a12.14 12.14 0 0 0 5.511 5.51l.932-1.256a1 1 0 0 1 1.431-.183l3.012 2.43a1 1 0 0 1 .203 1.335l-.888 1.324a3.5 3.5 0 0 1-4.331 1.247l-.68-.303Z"/><path fill="currentColor" d="M13.75 9a.75.75 0 0 0 .75.75h3.828a.75.75 0 0 0 0-1.5h-2.017l3.159-3.159a.75.75 0 1 0-1.061-1.06L15.25 7.188V5.172a.75.75 0 0 0-1.5 0V9Z"/>`),
		g.Group(children),
	)
}