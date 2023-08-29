package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleDownFilledCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13.866 18.499a1 1 0 0 1-1.732 0l-5.769-10A1 1 0 0 1 7.231 7H18.77a1 1 0 0 1 .866 1.5l-5.769 9.999Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M6.365 8.5A1 1 0 0 1 7.231 7H18.77a1 1 0 0 1 .866 1.5l-5.769 9.999a1 1 0 0 1-1.732 0l-5.769-10ZM13 13.997L15.306 10h-4.612L13 13.997Z" clip-rule="evenodd"/><path d="M4.293 5.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/><path fill-rule="evenodd" d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}