package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Headset(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTHeadset0"><g fill="none"><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M42 30v-5.538C42 14.265 33.941 6 24 6S6 14.265 6 24.462V30"/><path fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M34 32a4 4 0 0 1 4-4h4v14h-4a4 4 0 0 1-4-4v-6Z"/><path fill="#fff" d="M42 32h2a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2h-2v-6ZM6 32H4a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h2v-6Z"/><path fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M6 28h4a4 4 0 0 1 4 4v6a4 4 0 0 1-4 4H6V28Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTHeadset0)"/>`),
		g.Group(children),
	)
}