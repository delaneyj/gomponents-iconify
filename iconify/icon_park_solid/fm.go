package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSFm0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="22" r="4" fill="#fff"/><path fill="#fff" d="M21.524 33.054c1.238-1.405 3.714-1.405 4.952 0s0 10.307-.825 11.243c-.826.937-2.476.937-3.302 0c-.825-.936-2.063-9.838-.825-11.243Z"/><path d="M30.967 30.513A10.978 10.978 0 0 0 35 22c0-6.075-4.925-11-11-11s-11 4.925-11 11c0 3.431 1.571 6.496 4.033 8.513"/><path d="M31.926 38.166C37.893 35.235 42 29.096 42 22c0-9.941-8.059-18-18-18S6 12.059 6 22c0 7.097 4.107 13.234 10.074 16.166"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSFm0)"/>`),
		g.Group(children),
	)
}