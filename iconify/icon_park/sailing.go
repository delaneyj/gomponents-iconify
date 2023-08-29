package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sailing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M19 30H11L19 21"/><path stroke-linejoin="round" d="M39 30C39 17.0082 28.9937 4 19 4V30H39Z"/><path fill="#2F88FF" stroke-linejoin="round" d="M32.6512 41.5765L42 36L6 36L8 42L31.1144 42C31.6555 42 32.1865 41.8537 32.6512 41.5765Z"/><path stroke-linejoin="round" d="M19 30V36"/><path d="M29 21L41 21"/></g>`),
		g.Group(children),
	)
}