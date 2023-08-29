package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RemindDisable(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTRemindDisable0"><g fill="none"><path fill="#555" d="M24 4c-7.732 0-14 6.268-14 14v20h28V18c0-7.732-6.268-14-14-14Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M10 38V18c0-7.732 6.268-14 14-14s14 6.268 14 14v20M4 38h40"/><path fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 44a5 5 0 0 0 5-5v-1H19v1a5 5 0 0 0 5 5Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m19 17l10 10m0-10L19 27"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTRemindDisable0)"/>`),
		g.Group(children),
	)
}