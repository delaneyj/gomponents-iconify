package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTShieldAdd0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M6 8.256L24.009 3L42 8.256v10.778A26.316 26.316 0 0 1 24.003 44A26.32 26.32 0 0 1 6 19.029V8.256Z"/><path stroke-linecap="round" d="M23.95 14.95v16m-8-8h16"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTShieldAdd0)"/>`),
		g.Group(children),
	)
}