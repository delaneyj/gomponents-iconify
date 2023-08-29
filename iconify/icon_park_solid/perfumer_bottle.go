package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PerfumerBottle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSPerfumerBottle0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="38" height="24" x="5" y="17" fill="#fff" stroke="#fff" rx="2"/><path fill="#fff" stroke="#fff" d="M14 7h20v10H14z"/><path fill="#000" stroke="#000" d="M18 25h12v8H18z"/><path stroke="#000" d="M30 29h13M5 29h13"/><path stroke="#fff" d="M5 24v10m38-10v10"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSPerfumerBottle0)"/>`),
		g.Group(children),
	)
}