package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwoDimensionalCode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTTwoDimensionalCode0"><g fill="none" stroke="#fff" stroke-width="4"><path fill="#555" stroke-linejoin="round" d="M20 6H6v14h14V6Zm0 22H6v14h14V28ZM42 6H28v14h14V6Z"/><path stroke-linecap="round" d="M29 28v14m12-14v14"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTTwoDimensionalCode0)"/>`),
		g.Group(children),
	)
}