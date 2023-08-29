package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func At(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2C6.5 2 2 6.5 2 12s4.5 10 10 10c1.8 0 3.5-.5 5-1.3c.5-.3.6-.9.4-1.4c-.3-.5-.9-.6-1.4-.4c-3.8 2.2-8.7.9-10.9-2.9C2.9 12.2 4.2 7.3 8 5.1c3.8-2.2 8.7-.9 10.9 2.9c.7 1.2 1.1 2.6 1.1 4v.8c0 1-.8 1.8-1.8 1.8s-1.8-.8-1.8-1.8V8.5c0-.6-.4-1-1-1c-.5 0-.9.3-1 .8c-2-1.4-4.9-.9-6.3 1.1c-1.4 2-.9 4.9 1.1 6.3c1.9 1.3 4.4 1 5.9-.7c1.3 1.6 3.6 1.9 5.2.7c.9-.7 1.5-1.8 1.5-3V12C22 6.5 17.5 2 12 2zm0 12.5c-1.4 0-2.5-1.1-2.5-2.5s1.1-2.5 2.5-2.5s2.5 1.1 2.5 2.5s-1.1 2.5-2.5 2.5z"/>`),
		g.Group(children),
	)
}