package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpsideDownFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSUpsideDownFace0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M24 4C12.954 4 4 12.954 4 24s8.954 20 20 20s20-8.954 20-20S35.046 4 24 4Z"/><path stroke="#000" stroke-linecap="round" d="M17 30v-1m14 1v-1M17 17s2-4 7-4s7 4 7 4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSUpsideDownFace0)"/>`),
		g.Group(children),
	)
}