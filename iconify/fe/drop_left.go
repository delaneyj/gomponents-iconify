package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DropLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feDropLeft0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feDropLeft1" fill="currentColor"><path id="feDropLeft2" d="m16 5l-8 7l8 7z"/></g></g>`),
		g.Group(children),
	)
}