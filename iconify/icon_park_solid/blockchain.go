package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Blockchain(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBlockchain0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M40 30V15L27.5 7.969m-7 0L8 15v15m3 4.688L24 42l8-4.5l5-2.813"/><path fill="#fff" d="m21 18.75l-3 1.75v7l3 1.75L24 31l3-1.75l3-1.75v-7l-3-1.75L24 17l-3 1.75Z"/><path d="M24 17v-7m6 17l7 4m-19-4l-7 4"/><circle cx="24" cy="7" r="3" fill="#fff"/><circle cx="8" cy="33" r="3" fill="#fff"/><circle cx="40" cy="33" r="3" fill="#fff"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBlockchain0)"/>`),
		g.Group(children),
	)
}