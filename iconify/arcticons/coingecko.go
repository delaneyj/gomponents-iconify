package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Coingecko(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.051 37.228c4.519-9.359 1.828-21.471 9.117-23.378s12.056 1.346 15.87 2.468s6.978 1.716 7.682 5.27C40.729 26.693 24.13 35.71 27.664 45.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.355 26.972c3.196 2.523 6.496 1.685 13.796-1.373m-8.302-10.062c0-2.273-4.148-4.098-6.719-2.079"/><circle cx="21.047" cy="19.122" r="3.308" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="21.047" cy="19.122" r="1.29" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="34.841" cy="22.991" r=".775" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.375 22.99A18.404 18.404 0 0 0 25.033 5.628"/>`),
		g.Group(children),
	)
}