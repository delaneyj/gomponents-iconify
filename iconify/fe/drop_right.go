package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DropRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feDropRight0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feDropRight1" fill="currentColor"><path id="feDropRight2" d="M8 5v14l8-7z"/></g></g>`),
		g.Group(children),
	)
}