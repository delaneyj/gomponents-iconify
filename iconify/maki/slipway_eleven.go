package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlipwayEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M1.042 8l8 1.38V10h-8zm8-4l-1 1v.625L3.735 4.01l.077-.285l.86-.845H5.5a.35.35 0 0 0 0-.7h-.971a.348.348 0 0 0-.245.1l-1.03 1.011a.357.357 0 0 0-.092.158l-.086.314L1.042 3C1.095 4 1 5.51 3 5.884C3.983 6.07 8.095 7 8.095 7l1.947-2z" fill="currentColor"/>`),
		g.Group(children),
	)
}