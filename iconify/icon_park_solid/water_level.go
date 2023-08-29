package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterLevel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSWaterLevel0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M24 44c8.284 0 15-6.716 15-15C39 15 24 4 24 4S9 15 9 29c0 8.284 6.716 15 15 15Z" clip-rule="evenodd"/><path fill="#fff" d="M9 29c0 8.284 6.716 15 15 15c8.284 0 15-6.716 15-15c0 0-9 3-15 0S9 29 9 29Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSWaterLevel0)"/>`),
		g.Group(children),
	)
}