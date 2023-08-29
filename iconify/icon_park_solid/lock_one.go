package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSLockOne0"><g fill="none" stroke-width="4"><circle cx="24" cy="30" r="14" fill="#fff" stroke="#fff"/><path stroke="#fff" stroke-linejoin="round" d="M31 18v-7a7 7 0 1 0-14 0v7"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M24 26v8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSLockOne0)"/>`),
		g.Group(children),
	)
}