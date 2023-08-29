package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RedEnvelope(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTRedEnvelope0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M39 4H9v6l15 4l15-4V4Z"/><path d="M39 17v27H9V17"/><path d="m19 19l5 6l5-6M18 31h12m-12-6h12m-6 0v12"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTRedEnvelope0)"/>`),
		g.Group(children),
	)
}