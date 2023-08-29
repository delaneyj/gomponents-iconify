package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PerfumerBottle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTPerfumerBottle0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="38" height="24" x="5" y="17" fill="#555" rx="2"/><path fill="#555" d="M14 7h20v10H14zm4 18h12v8H18z"/><path d="M30 29h13M5 29h13M5 24v10m38-10v10"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTPerfumerBottle0)"/>`),
		g.Group(children),
	)
}