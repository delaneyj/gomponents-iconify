package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Certificate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M26 36H6C4.89543 36 4 35.1046 4 34V8C4 6.89543 4.89543 6 6 6H42C43.1046 6 44 6.89543 44 8V34C44 35.1046 43.1046 36 42 36H34"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 14H36"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 21H18"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 28H16"/><path fill="#2F88FF" d="M30 33C33.3137 33 36 30.3137 36 27C36 23.6863 33.3137 21 30 21C26.6863 21 24 23.6863 24 27C24 30.3137 26.6863 33 30 33Z"/><path fill="#2F88FF" stroke-linecap="round" stroke-linejoin="round" d="M30 40L34 42V31.4722C34 31.4722 32.8594 33 30 33C27.1406 33 26 31.5 26 31.5V42L30 40Z"/></g>`),
		g.Group(children),
	)
}