package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignLeftOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSAlignLeftOne0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><path fill="#fff" stroke-linejoin="round" d="M16 6h16v6H16V6Z"/><path d="M6 42V6"/><path fill="#fff" stroke-linejoin="round" d="M16 21h20v6H16v-6Zm0 15h26v6H16v-6Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSAlignLeftOne0)"/>`),
		g.Group(children),
	)
}