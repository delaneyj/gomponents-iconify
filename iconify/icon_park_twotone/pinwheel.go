package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pinwheel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTPinwheel0"><path fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M21 4v17H11L21 4Zm6 40V27h10L27 44Zm0-33l17 10H27V11Zm-6 26L4 27h17v10Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTPinwheel0)"/>`),
		g.Group(children),
	)
}