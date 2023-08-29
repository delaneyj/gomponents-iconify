package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Maildotcom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.017-.009H0V15.66c0 1.406.96 2.571 2.246 2.914L24 24.008V5.992c.017-3.308-2.674-6-5.983-6zm3 15.669H18V8.786c0-.669-.223-2.229-2.212-2.229c-1.32 0-2.28.909-2.28 2.229v6.874h-3.017V8.786c0-.669-.205-2.229-2.194-2.229c-1.354 0-2.28.909-2.28 2.229v6.874H3V3.609h5.297c1.594 0 2.897.634 3.737 1.662c.892-1.028 2.212-1.662 3.737-1.662C19.063 3.609 21 5.786 21 8.854l.017 6.806z"/>`),
		g.Group(children),
	)
}