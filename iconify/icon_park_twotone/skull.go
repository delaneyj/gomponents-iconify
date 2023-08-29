package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Skull(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTSkull0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="m11.263 44l3.364-8.315C10.031 32.593 7 27.293 7 21.273C7 11.733 14.611 4 24 4s17 7.733 17 17.273c0 6.02-3.031 11.32-7.627 14.412L36.737 44H11.263Z"/><path stroke-linecap="round" d="M20 38v6m8-6v6"/><path fill="#555" d="M17 23a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm14 0a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/><path stroke-linecap="round" d="M32 44h-8m0 0h-8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTSkull0)"/>`),
		g.Group(children),
	)
}