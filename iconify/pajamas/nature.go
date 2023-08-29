package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nature(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.98 2.5A6.37 6.37 0 0 0 15 2V1h-1.75a6.003 6.003 0 0 0-5.761 4.32A5.99 5.99 0 0 0 2.75 3H1v1a6 6 0 0 0 6 6h.25v4.25a.75.75 0 0 0 1.5 0V8H9a6 6 0 0 0 5.98-5.5Zm-6.203 4a4.5 4.5 0 0 1 4.473-4h.223A4.5 4.5 0 0 1 9 6.5h-.223Zm-6.027-2a4.5 4.5 0 0 1 4.473 4H7a4.5 4.5 0 0 1-4.473-4h.223Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}