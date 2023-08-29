package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FirstAidKitSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path id="clarityFirstAidKitSolid0" fill="currentColor" d="M32 6h-8.09V4.5a2.5 2.5 0 0 0-2.5-2.5h-7a2.5 2.5 0 0 0-2.5 2.5V6H4a2 2 0 0 0-2 2v20a2 2 0 0 0 2 2h28a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2ZM13.91 4.5a.5.5 0 0 1 .5-.5h7a.5.5 0 0 1 .5.5V6h-8Zm10.73 15.4a.5.5 0 0 1-.5.5h-3.5v3.5a.5.5 0 0 1-.5.5h-3.4a.5.5 0 0 1-.5-.5v-3.5h-3.5a.5.5 0 0 1-.5-.5v-3.4a.5.5 0 0 1 .5-.5h3.5v-3.5a.5.5 0 0 1 .5-.5h3.4a.5.5 0 0 1 .5.5V16h3.5a.5.5 0 0 1 .5.5Z"/>`),
		g.Group(children),
	)
}