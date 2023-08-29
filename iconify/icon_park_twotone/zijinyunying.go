package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zijinyunying(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTZijinyunying0"><g fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M34 6H14L3 24l11 18h20l11-18L34 6Z"/><path d="m15 29l9-14l9 14H15Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTZijinyunying0)"/>`),
		g.Group(children),
	)
}