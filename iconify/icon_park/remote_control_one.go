package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RemoteControlOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="#000" stroke-width="4" d="M11 5.44578C11 4.6473 11.6473 4 12.4458 4H35.5542C36.3527 4 37 4.6473 37 5.44578V31C37 38.1797 31.1797 44 24 44V44C16.8203 44 11 38.1797 11 31V5.44578Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M17 16H20"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M28 16H31"/><circle cx="17" cy="10" r="2" fill="#000"/><circle cx="24" cy="31" r="7" fill="#2F88FF" stroke="#000" stroke-width="4"/></g>`),
		g.Group(children),
	)
}