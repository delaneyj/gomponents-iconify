package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Robot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTRobot0"><g fill="none" stroke="#fff" stroke-width="4"><rect width="30" height="26" x="9" y="17" stroke-linecap="round" stroke-linejoin="round" rx="2"/><path stroke-linecap="round" stroke-linejoin="round" d="m33 9l-5 8M15 9l5 8"/><circle cx="34" cy="7" r="2"/><circle cx="14" cy="7" r="2"/><rect width="16" height="8" x="16" y="24" fill="#555" rx="4"/><path stroke-linecap="round" stroke-linejoin="round" d="M9 24H4v10h5m30-10h5v10h-5"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTRobot0)"/>`),
		g.Group(children),
	)
}