package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Airplay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTAirplay0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M12 35.014H4V8.013a2 2 0 0 1 2-2h36a2 2 0 0 1 2 2v27h-8"/><path fill="#555" d="M24 32L14 42h20L24 32Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTAirplay0)"/>`),
		g.Group(children),
	)
}