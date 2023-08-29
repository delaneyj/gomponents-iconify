package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bolt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feBolt0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feBolt1" fill="currentColor"><path id="feBolt2" d="M18 2h-8L6 13h4l-2 9l9-13h-4.995z"/></g></g>`),
		g.Group(children),
	)
}