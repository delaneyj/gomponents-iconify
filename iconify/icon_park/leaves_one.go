package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeavesOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M37 23.8788C37 30.5731 31.1797 36 24 36C16.8203 36 11 30.5731 11 23.8788C11 17.1844 24 4 24 4C24 4 37 17.1844 37 23.8788Z"/><path stroke="#fff" stroke-linecap="round" d="M24 4V36"/><path stroke="#000" stroke-linecap="round" d="M24 36V44"/><path stroke="#000" d="M37 23.8789C37 30.5733 31.1797 36.0001 24 36.0001C16.8203 36.0001 11 30.5733 11 23.8789"/><path stroke="#000" d="M37 23.8788C37 17.1844 24 4 24 4C24 4 11 17.1844 11 23.8788"/></g>`),
		g.Group(children),
	)
}