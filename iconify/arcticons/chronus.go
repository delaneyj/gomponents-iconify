package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chronus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="23.471" cy="24" r="3.162" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.471 20.838V4.5m0 39V32.984m2.728-7.394l14.132 8.236m-19.59-8.235l-14.13 8.235m11.602 4.416h10.516m5.362-8.051a6.532 6.532 0 1 0 .766-13.017q-.2 0-.396.011v.002a11.842 11.842 0 0 0-10.99-6.944a13.76 13.76 0 0 0-11.887 20.685"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.471 15.04a8.962 8.962 0 0 0-7.742 13.472m18.732-11.325a4.59 4.59 0 0 0-3.36 1.194"/>`),
		g.Group(children),
	)
}