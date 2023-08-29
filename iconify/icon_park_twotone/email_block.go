package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmailBlock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTEmailBlock0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M44 35V9H4v28h22"/><circle cx="35" cy="35" r="9" fill="#555"/><path stroke-linecap="round" stroke-linejoin="round" d="m37 33l-4 4M4 9l20 13L44 9"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTEmailBlock0)"/>`),
		g.Group(children),
	)
}