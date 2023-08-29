package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scale(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feScale0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feScale1" fill="currentColor"><path id="feScale2" d="M10.874 17.998A4.002 4.002 0 0 1 3 17a4.002 4.002 0 0 1 3.002-3.874A5.001 5.001 0 0 1 9.03 8.404a6 6 0 1 1 6.567 6.567a5.001 5.001 0 0 1-4.723 3.027Zm.252-9.996a5 5 0 0 1 4.872 4.872A4.002 4.002 0 0 0 15 5a4.002 4.002 0 0 0-3.874 3.002Zm-.253 7.995a3 3 0 1 0-2.87-2.87a4.007 4.007 0 0 1 2.87 2.87ZM7 19a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/></g></g>`),
		g.Group(children),
	)
}