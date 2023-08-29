package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoneyBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M192 299v-22h-43v-42h86v-22h-64q-9 0-15.5-6t-6.5-15v-64q0-9 6.5-15t15.5-6h21V85h43v22h42v42h-85v22h64q9 0 15 6t6 15v64q0 9-6 15t-15 6h-21v22h-43zM384 21q18 0 30.5 12.5T427 64v256q0 18-12.5 30.5T384 363H43q-18 0-30.5-12.5T0 320V64q0-18 12.5-30.5T43 21h341zm0 299V64H43v256h341z"/>`),
		g.Group(children),
	)
}