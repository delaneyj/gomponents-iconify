package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoybeanMilkMaker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path stroke="#000" d="M35 10H40.0644C40.5909 10 41.0272 10.4082 41.0622 10.9335L41.9289 23.9335C41.9674 24.5107 41.5096 25 40.9311 25H32"/><path stroke="#000" d="M7 5H35L31 31H15L12 10L7 5Z"/><path fill="#2F88FF" stroke="#000" d="M15 31H31L35 43H11L15 31Z"/><path stroke="#fff" d="M21 37H25"/></g>`),
		g.Group(children),
	)
}