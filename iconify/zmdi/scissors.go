package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scissors(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m163 123l264 264v21h-64L213 259l-50 50q8 17 8 35q0 35-25 60t-60.5 25T25 404T0 344t25-60t60-25q19 0 35 7l51-50l-51-50q-16 7-35 7q-35 0-60-25T0 88t25-60T85.5 3T146 28t25 60q0 18-8 35zm-77.5 8q17.5 0 30-12.5T128 88t-12.5-30.5t-30-12.5t-30 12.5T43 88t12.5 30.5t30 12.5zm0 256q17.5 0 30-12.5T128 344t-12.5-30.5t-30-12.5t-30 12.5T43 344t12.5 30.5t30 12.5zm128-160q10.5 0 10.5-11t-10.5-11t-10.5 11t10.5 11zM363 24h64v21L277 195l-42-43z"/>`),
		g.Group(children),
	)
}