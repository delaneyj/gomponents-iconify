package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MassageChair(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="38" height="38" x="5" y="5" fill="#2F88FF" stroke="#000" rx="3"/><path stroke="#fff" d="M18 24V15.6522C18 14.4348 19.2 12 24 12C28.8 12 30 14.4348 30 15.6522V24"/><path stroke="#fff" d="M16 24V30H32V24"/><path stroke="#fff" d="M12 20V24H36V20"/><path stroke="#fff" d="M18 36L30 36"/><path stroke="#fff" d="M24 30V36"/></g>`),
		g.Group(children),
	)
}