package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Code(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feCode0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feCode1" fill="currentColor"><path id="feCode2" d="m21.004 11.657l-5.657 5.657l-1.414-1.415l4.242-4.242l-4.242-4.243L15.347 6l5.657 5.657Zm-15.176 0l4.243 4.242l-1.414 1.415L3 11.657L8.657 6l1.414 1.414l-4.243 4.243Z"/></g></g>`),
		g.Group(children),
	)
}