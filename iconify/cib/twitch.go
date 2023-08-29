package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Twitch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M2.787 0L.698 5.568v22.255H8.35V32h4.177l4.167-4.177h6.26l8.349-8.344V0zm2.781 2.781h22.953v15.301l-4.871 4.871h-7.651l-4.172 4.172v-4.172h-6.26zm7.651 13.916H16V8.349h-2.781zm7.645 0h2.781V8.349h-2.781z"/>`),
		g.Group(children),
	)
}