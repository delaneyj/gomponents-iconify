package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessageSuccess(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M25.5 36H21L11 41V36H4V6H44V17"/><path d="M12 14H15L18 14"/><path d="M12 20H18L24 20"/><path d="M29 30L35 35L44 24"/></g>`),
		g.Group(children),
	)
}