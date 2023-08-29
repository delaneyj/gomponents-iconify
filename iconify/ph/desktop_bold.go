package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DesktopBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 36H48a28 28 0 0 0-28 28v108a28 28 0 0 0 28 28h68v12H96a12 12 0 0 0 0 24h64a12 12 0 0 0 0-24h-20v-12h68a28 28 0 0 0 28-28V64a28 28 0 0 0-28-28ZM48 60h160a4 4 0 0 1 4 4v72H44V64a4 4 0 0 1 4-4Zm160 116H48a4 4 0 0 1-4-4v-12h168v12a4 4 0 0 1-4 4Z"/>`),
		g.Group(children),
	)
}