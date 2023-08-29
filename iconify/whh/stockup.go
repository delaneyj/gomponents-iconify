package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stockup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm0-768q0-53-37.5-90.5t-90.5-37.5h-512q-53 0-90.5 37.5t-37.5 90.5v512q0 53 37.5 90.5t90.5 37.5h512q53 0 90.5-37.5t37.5-90.5V256zm-111 318l-60-60l-188 188q-6 11-11 16q-18 18-43.5 18t-43.5-18l-55-55l-88 87q-18 18-43.5 18t-43.5-18t-18-43.5t18-43.5l122-122q7-10 7-11q18-18 43.5-18t43.5 18l54 55l158-158l-61-61v-6q0-14 .5-21t6-13t16.5-6h186q20 0 33.5 13.5t13.5 32.5v185q0 11-6 16.5t-13 6t-21 .5h-7z"/>`),
		g.Group(children),
	)
}