package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Speaker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M256 3q18 0 30.5 12.5T299 45v342q0 17-12.5 29.5T256 429H43q-18 0-30.5-12.5T0 387V45q0-17 12.5-29.5T43 3h213zM149.5 45q-17.5 0-30 12.5T107 88t12.5 30.5t30 12.5t30-12.5T192 88t-12.5-30.5t-30-12.5zm-.5 342q44 0 75.5-31.5T256 280t-31.5-75.5T149 173t-75 31.5T43 280t31 75.5t75 31.5zm.5-171q26.5 0 45 18.5T213 280t-18.5 45.5t-45 18.5t-45.5-18.5T85 280t19-45.5t45.5-18.5z"/>`),
		g.Group(children),
	)
}