package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Parking(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSParking0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M24 44s16-12 16-25c0-8.284-7.163-15-16-15S8 10.716 8 19c0 13 16 25 16 25Z"/><path stroke="#000" stroke-linecap="round" d="M21 14v16"/><path fill="#000" stroke="#000" d="M21 14h6a4 4 0 0 1 0 8h-6v-8Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSParking0)"/>`),
		g.Group(children),
	)
}