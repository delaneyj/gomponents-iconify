package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mascara(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTMascara0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M28 16h12v28H28z"/><path fill="#555" stroke-linejoin="round" d="M8 28h12v16H8z"/><path stroke-linejoin="round" d="M14 4v24"/><path d="M20 36H8"/><path stroke-linejoin="round" d="M20 32v8M8 32v8M18 9h-8m10 6H8m10 6h-8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTMascara0)"/>`),
		g.Group(children),
	)
}