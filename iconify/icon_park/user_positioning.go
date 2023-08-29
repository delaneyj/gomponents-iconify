package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserPositioning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="16" r="6" fill="#2F88FF"/><path d="M36 36C36 29.3726 30.6274 24 24 24C17.3726 24 12 29.3726 12 36"/><path d="M36 4H44V12"/><path d="M12 4H4V12"/><path d="M36 44H44V36"/><path d="M12 44H4V36"/></g>`),
		g.Group(children),
	)
}