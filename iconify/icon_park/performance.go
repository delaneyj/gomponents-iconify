package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Performance(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M24 4C12.9543 4 4 12.9543 4 24C4 35.0457 12.9543 44 24 44C35.0457 44 44 35.0457 44 24C44 21.2883 43.4603 18.7026 42.4825 16.3446C42.2308 15.7376 41.9501 15.1457 41.6421 14.5707"/><path fill="#2F88FF" d="M35 10C37.2091 10 39 8.65685 39 7C39 5.34315 37.2091 4 35 4C32.7909 4 31 5.34315 31 7C31 8.65685 32.7909 10 35 10Z"/><path fill="#2F88FF" d="M24 31C27.866 31 31 27.866 31 24C31 20.134 27.866 17 24 17C20.134 17 17 20.134 17 24C17 27.866 20.134 31 24 31Z"/><path stroke-linecap="round" d="M31 6.5V24"/></g>`),
		g.Group(children),
	)
}