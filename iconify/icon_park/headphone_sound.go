package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeadphoneSound(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path fill="#2F88FF" d="M4 28C4 26.8954 4.89543 26 6 26H10V38H6C4.89543 38 4 37.1046 4 36V28Z"/><path fill="#2F88FF" d="M38 26H42C43.1046 26 44 26.8954 44 28V36C44 37.1046 43.1046 38 42 38H38V26Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M10 36V24C10 16.268 16.268 10 24 10C31.732 10 38 16.268 38 24V36M10 26H6C4.89543 26 4 26.8954 4 28V36C4 37.1046 4.89543 38 6 38H10V26ZM38 26H42C43.1046 26 44 26.8954 44 28V36C44 37.1046 43.1046 38 42 38H38V26Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M16 32H20L22 26L26 38L28 32H32"/></g>`),
		g.Group(children),
	)
}