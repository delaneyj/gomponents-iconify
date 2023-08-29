package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageCircleFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M37.096 39.116A19.923 19.923 0 0 1 24 44a19.923 19.923 0 0 1-13.096-4.884l12.39-12.325a1 1 0 0 1 1.411 0l12.391 12.325Zm1.784-1.752L26.468 25.02a3.5 3.5 0 0 0-4.936 0L9.12 37.364A19.926 19.926 0 0 1 4 24C4 12.954 12.954 4 24 4s20 8.954 20 20c0 5.137-1.937 9.822-5.12 13.364ZM35 17a4 4 0 1 0-8 0a4 4 0 0 0 8 0Z"/>`),
		g.Group(children),
	)
}