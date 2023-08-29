package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmailLock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSEmailLock0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M44 24V9H4v30h20"/><path d="m4 9l20 15L44 9"/><path fill="#fff" d="M31 33h12v8H31z"/><path d="M40 33v-3a3 3 0 1 0-6 0v3"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSEmailLock0)"/>`),
		g.Group(children),
	)
}