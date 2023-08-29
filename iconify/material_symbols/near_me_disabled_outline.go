package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NearMeDisabledOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12.9 21l-2.85-7.05L3 11.1V9.7l4.875-1.825L2.8 2.8l1.425-1.425l18.4 18.4L21.2 21.2l-5.075-5.075L14.3 21h-1.4Zm4.775-9.025L16.1 10.4l1.5-4l-4 1.5l-1.575-1.575L21 3l-3.325 8.975ZM13.55 17.3l1.025-2.725l-5.15-5.15L6.7 10.45l4.9 1.95l1.95 4.9Zm1.3-8.15ZM12 12Z"/>`),
		g.Group(children),
	)
}