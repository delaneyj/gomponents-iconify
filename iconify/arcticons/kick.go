package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kick(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.75 21.842v-4.25H41V5H28.25v4.25H24v4.25h-4.25V5H7v38h12.75v-8.5H24v4.25h4.25V43H41V30.342h-4.25v-4.25H32.5v-4.25h4.25z"/>`),
		g.Group(children),
	)
}