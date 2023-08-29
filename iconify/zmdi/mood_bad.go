package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoodBad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M213.5 3q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3zm0 384q70.5 0 120.5-50t50-121t-50-121t-120.5-50T93 95T43 216t50 121t120.5 50zM288 195q-13 0-22.5-9.5t-9.5-23t9.5-22.5t22.5-9t22.5 9t9.5 22.5t-9.5 23T288 195zm-149.5 0q-13.5 0-22.5-9.5t-9-23t9-22.5t22.5-9t23 9t9.5 22.5t-9.5 23t-23 9.5zm75 64q36.5 0 66 20.5T322 333H104q13-33 43-53.5t66.5-20.5z"/>`),
		g.Group(children),
	)
}