package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lamusica(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.081 22.632c-.317-7.247 7.998-3.955 11.607-11.174a32.271 32.271 0 0 1-7.978 3.666c-4.261 1.448-6.534 2.206-3.629 7.508Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.53 17.588c.334-6.205 9.654-.543 16.584-13.088c-8.69 7.738-20.789 5.428-16.583 13.088Zm2.551 5.044l5.628 10.118M18.53 17.588l1.754 3.394"/><ellipse cx="20.179" cy="36.929" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="7.753" ry="6.02" transform="rotate(-32.619 20.18 36.93)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24.052 37.425l-5.809-3.354v6.707l5.809-3.353z"/>`),
		g.Group(children),
	)
}