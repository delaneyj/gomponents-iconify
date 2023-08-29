package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Routeralt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.615 704h-128l182-660q4-15 16.5-30.5t25.5-12.5q15 3 24.5 14t7.5 23zM.615 38q-1-12 8-23t24-14q13-3 25.5 12.5t16.5 30.5l182 660h-128zm128 730h768q53 0 90.5 37.5t37.5 90.5t-37.5 90.5t-90.5 37.5h-768q-53 0-90.5-37.5T.615 896t37.5-90.5t90.5-37.5zm639.5 192q26.5 0 45.5-18.5t19-45t-19-45.5t-45.5-19t-45 19t-18.5 45.5t18.5 45t45 18.5zm-255.5 0q26 0 45-18.5t19-45t-19-45.5t-45.5-19t-45 19t-18.5 45.5t19 45t45 18.5zm-255.5 0q26.5 0 45-18.5t18.5-45t-18.5-45.5t-45-19t-45.5 19t-19 45.5t19 45t45.5 18.5z"/>`),
		g.Group(children),
	)
}