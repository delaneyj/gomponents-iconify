package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Facebook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTFacebook0"><path fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="3.8" d="M36 12.6h-6.013c-1.086 0-1.967.88-1.967 1.967v6.9H36l-1.169 7.597h-6.81V43h-8.776V29.064H12v-7.597h7.151l.094-7.21l-.013-1.31A7.868 7.868 0 0 1 27.099 5H36v7.6Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTFacebook0)"/>`),
		g.Group(children),
	)
}