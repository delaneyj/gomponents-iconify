package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSLock0"><g fill="none" stroke-linejoin="round" stroke-width="4"><rect width="36" height="22" x="6" y="22" fill="#fff" stroke="#fff" rx="2"/><path stroke="#fff" stroke-linecap="round" d="M14 22v-8c0-5.523 4.477-10 10-10s10 4.477 10 10v8"/><path stroke="#000" stroke-linecap="round" d="M24 30v6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSLock0)"/>`),
		g.Group(children),
	)
}