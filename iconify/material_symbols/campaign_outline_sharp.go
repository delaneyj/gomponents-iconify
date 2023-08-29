package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CampaignOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 13v-2h4v2h-4Zm1.2 7L16 17.6l1.2-1.6l3.2 2.4l-1.2 1.6Zm-2-12L16 6.4L19.2 4l1.2 1.6L17.2 8ZM5 19v-4H2V9h6l5-3v12l-5-3H7v4H5Zm2.5-7Zm6.5 3.35v-6.7q.675.6 1.088 1.463T15.5 12q0 1.025-.413 1.888T14 15.35ZM4 13h4.55L11 14.45v-4.9L8.55 11H4v2Z"/>`),
		g.Group(children),
	)
}