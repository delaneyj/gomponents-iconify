package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnlockOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTUnlockOne0"><g fill="none" stroke="#fff" stroke-width="4"><circle cx="24" cy="30" r="14" fill="#555"/><path stroke-linecap="round" stroke-linejoin="round" d="M31 12v-1a7 7 0 0 0-7-7v0a7 7 0 0 0-7 7v6m7 9v8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTUnlockOne0)"/>`),
		g.Group(children),
	)
}