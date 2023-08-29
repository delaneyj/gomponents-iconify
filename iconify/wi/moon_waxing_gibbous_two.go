package wi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoonWaxingGibbousTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 30 30"),
		g.Raw(`<path fill="currentColor" d="M12.85 14.44c0 4.77.71 8.52 2.14 11.26c2.04 0 3.93-.5 5.65-1.51s3.1-2.37 4.1-4.1s1.51-3.61 1.51-5.65s-.5-3.92-1.51-5.65s-2.37-3.09-4.1-4.09s-3.61-1.51-5.65-1.51c-1.42 3.42-2.14 7.17-2.14 11.25z"/>`),
		g.Group(children),
	)
}