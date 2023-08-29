package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MouseOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><rect width="24" height="32" x="12" y="12" rx="12"/><path fill="#2F88FF" d="M12 24C12 17.3726 17.3726 12 24 12C30.6274 12 36 17.3726 36 24V25H12V24Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M24 25C24 25 24 16 24 12C24 8 25.5 4 31 4C37 4 39 9 39 13"/></g>`),
		g.Group(children),
	)
}