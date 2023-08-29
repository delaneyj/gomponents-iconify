package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="22" r="4"/><path d="M21.524 33.054c1.238-1.405 3.714-1.405 4.952 0s0 10.307-.825 11.243c-.826.937-2.476.937-3.302 0c-.825-.936-2.063-9.838-.825-11.243Zm9.443-2.541A10.978 10.978 0 0 0 35 22c0-6.075-4.925-11-11-11s-11 4.925-11 11c0 3.431 1.571 6.496 4.033 8.513"/><path d="M31.926 38.166C37.893 35.235 42 29.096 42 22c0-9.941-8.059-18-18-18S6 12.059 6 22c0 7.097 4.107 13.234 10.074 16.166"/></g>`),
		g.Group(children),
	)
}