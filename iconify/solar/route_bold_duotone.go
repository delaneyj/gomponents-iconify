package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RouteBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M12 2c-1.716 0-3.096 1.38-5.858 4.142C3.381 8.904 2 10.284 2 12c0 1.716 1.38 3.096 4.142 5.858C8.904 20.619 10.284 22 12 22c1.716 0 3.096-1.38 5.858-4.142C20.619 15.096 22 13.716 22 12c0-1.716-1.38-3.096-4.142-5.858C15.096 3.381 13.716 2 12 2Z" opacity=".5"/><path d="M12.786 8.487a.75.75 0 0 1 1.06-.034l2.667 2.5a.75.75 0 0 1 0 1.094l-2.667 2.5a.75.75 0 0 1-1.026-1.094l1.283-1.203h-3.436c-.334 0-.845.1-1.247.372c-.363.245-.67.643-.67 1.378a.75.75 0 0 1-1.5 0c0-1.265.582-2.117 1.33-2.622c.709-.478 1.532-.628 2.087-.628h3.436L12.82 9.547a.75.75 0 0 1-.034-1.06Z"/></g>`),
		g.Group(children),
	)
}