package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AreaChartOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M.5 14.5H0v.5h.5v-.5Zm6-9l.4-.3a.5.5 0 0 0-.816.023L6.5 5.5Zm3 4l-.4.3a.5.5 0 0 0 .807-.01L9.5 9.5ZM0 0v14.5h1V0H0Zm.5 15H15v-1H.5v1Zm2.416-3.223l4-6l-.832-.554l-4 6l.832.554ZM6.1 5.8l3 4l.8-.6l-3-4l-.8.6Zm3.807 3.99l5-7l-.814-.58l-5 7l.814.58Z"/>`),
		g.Group(children),
	)
}