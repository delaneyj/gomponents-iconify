package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DroneOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTDroneOne0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M29 18v-2a5 5 0 0 0-5-5v0a5 5 0 0 0-5 5v2"/><path fill="#555" stroke-linecap="round" d="M17 18h14l-2.154 7h-9.692L17 18Z"/><path fill="#555" d="M5 22h7v7H5zm31 0h7v7h-7z"/><path stroke-linecap="round" d="M16 8H4m26 25l4 7m-16-7l-4 7M44 8H32"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTDroneOne0)"/>`),
		g.Group(children),
	)
}