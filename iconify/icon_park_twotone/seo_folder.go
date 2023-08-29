package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SeoFolder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTSeoFolder0"><g fill="none" stroke="#fff" stroke-width="4"><path fill="#555" stroke-linejoin="round" d="M5 8a2 2 0 0 1 2-2h12l5 6h17a2 2 0 0 1 2 2v26a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V8Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m14 22l5 5l-5 5"/><path stroke-linecap="round" d="M26 32h8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTSeoFolder0)"/>`),
		g.Group(children),
	)
}