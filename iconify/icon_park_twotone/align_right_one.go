package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignRightOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTAlignRightOne0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><path d="M42 42V6"/><path fill="#555" stroke-linejoin="round" d="M16 6h16v6H16V6Zm-4 15h20v6H12v-6ZM6 36h26v6H6v-6Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTAlignRightOne0)"/>`),
		g.Group(children),
	)
}