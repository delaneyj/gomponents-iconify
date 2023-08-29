package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DesktopMac(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M427 3q17 0 29.5 12.5T469 45v256q0 18-12.5 30.5T427 344H277l43 64v21H149v-21l43-64H43q-18 0-30.5-12.5T0 301V45q0-17 12.5-29.5T43 3h384zm0 256V45H43v214h384z"/>`),
		g.Group(children),
	)
}