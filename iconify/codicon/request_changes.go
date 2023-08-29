package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RequestChanges(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m10.71 1.29l3 3L14 5v9l-1 1H4l-1-1V2l1-1h6l.71.29ZM4 14h9V5l-3-3H4v12Zm4-8H6v1h2v2h1V7h2V6H9V4H8v2Zm-2 5h5v1H6v-1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}