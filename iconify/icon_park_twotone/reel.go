package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Reel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTReel0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M31 10L14 20m20-3L14 29m20-3L15 38"/><path fill="#555" stroke-linecap="round" stroke-linejoin="round" d="M9 4h30v6H9zm0 34h30v6H9z"/><path stroke-linecap="round" stroke-linejoin="round" d="M14 10v28"/><path d="M34 10v28"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTReel0)"/>`),
		g.Group(children),
	)
}