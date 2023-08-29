package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileProtectionOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M40 23V14L31 4H10C8.89543 4 8 4.89543 8 6V42C8 43.1046 8.89543 44 10 44H22"/><path fill="#2F88FF" d="M27 31.2C27 30.1333 34 28 34 28C34 28 41 30.1333 41 31.2C41 39.7333 34 44 34 44C34 44 27 39.7333 27 31.2Z"/><path d="M30 4V14H40"/></g>`),
		g.Group(children),
	)
}