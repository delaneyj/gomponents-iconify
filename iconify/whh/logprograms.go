package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Logprograms(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-768-192q0 27 19 45.5t45 18.5h640q26 0 45-18.5t19-45.5V640h-160q-17 0-26-15l-134-134v246q0 13-9 22t-22 9h-1q-22 0-30-22l-165-219l-99 98q-9 15-26 15h-96v192zm704-704h-640q-26 0-45 19t-19 45v384h85l112-112q2-5 4-7q9-9 23-9t23 9q5 5 7 13l130 173V415q0-13 9-22t22-9h1q17 0 26 15l177 177h149V192q0-26-19-45t-45-19z"/>`),
		g.Group(children),
	)
}