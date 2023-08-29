package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnlockOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSUnlockOne0"><g fill="none" stroke-width="4"><circle cx="24" cy="30" r="14" fill="#fff" stroke="#fff"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M31 12v-1a7 7 0 0 0-7-7v0a7 7 0 0 0-7 7v6"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M24 26v8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSUnlockOne0)"/>`),
		g.Group(children),
	)
}