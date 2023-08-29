package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func XiaoduHome(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M4 20L24 6L44 20V42H4V20Z"/><path stroke="#fff" d="M12.6865 26.6863C14.1723 25.2006 15.9361 24.022 17.8773 23.2179C19.8185 22.4139 21.8991 22 24.0002 22C26.1014 22 28.182 22.4139 30.1232 23.2179C32.0644 24.022 33.8282 25.2006 35.314 26.6863"/><path stroke="#fff" d="M18.3428 32.3431C19.0856 31.6003 19.9676 31.011 20.9382 30.609C21.9088 30.2069 22.9491 30 23.9996 30C25.0502 30 26.0905 30.2069 27.0611 30.609C28.0317 31.011 28.9136 31.6003 29.6565 32.3431"/></g>`),
		g.Group(children),
	)
}