package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadarTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSRadarTwo0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M24 20V4C12.954 4 4 12.954 4 24s8.954 20 20 20s20-8.954 20-20"/><path fill="#fff" fill-rule="evenodd" d="M24 28a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z" clip-rule="evenodd"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSRadarTwo0)"/>`),
		g.Group(children),
	)
}