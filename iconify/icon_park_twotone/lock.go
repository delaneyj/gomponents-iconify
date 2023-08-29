package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTLock0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><rect width="36" height="22" x="6" y="22" fill="#555" rx="2"/><path stroke-linecap="round" d="M14 22v-8c0-5.523 4.477-10 10-10s10 4.477 10 10v8m-10 8v6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTLock0)"/>`),
		g.Group(children),
	)
}