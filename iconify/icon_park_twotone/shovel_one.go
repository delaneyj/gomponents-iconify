package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShovelOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTShovelOne0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><path fill="#555" stroke-linejoin="round" d="M11 4h26l-2.833 20H13.833L11 4Z"/><path d="M21 11v6m6-6v6m-3 7v20"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTShovelOne0)"/>`),
		g.Group(children),
	)
}