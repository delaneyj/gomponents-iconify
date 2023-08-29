package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GuideBoard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSGuideBoard0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M24 4v37"/><path fill="#fff" d="M24 8h15.545L42 12l-2.455 4H24V8Zm0 14H8.455L6 26l2.455 4H24v-8Z"/><path stroke-linecap="round" d="M16 42h16"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSGuideBoard0)"/>`),
		g.Group(children),
	)
}