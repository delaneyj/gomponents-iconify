package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FarlightEightyFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m31.071 13.845l6.006-6.4l-23.436-.169l-.489 1.986l-8.652.532l10.313 2.699l.017-1.27l16.241 2.622zm-15.384 3.843l-8.545-2.002l11.572 20.38l1.964-.569l4.787 7.227l-2.82-10.281l-1.107.62l-5.851-15.375zm11.02 11.401l2.539 8.401L41.11 17.279l-1.475-1.417L43.5 8.103l-7.494 7.582l1.092.649l-10.391 12.755zM22.911 25l6.385-6.248l-8.603-2.405L22.911 25z"/>`),
		g.Group(children),
	)
}