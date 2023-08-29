package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stockdown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm0-768q0-53-37.5-90.5t-90.5-37.5h-512q-53 0-90.5 37.5t-37.5 90.5v512q0 53 37.5 90.5t90.5 37.5h512q53 0 90.5-37.5t37.5-90.5V256zm-111 448h-186q-16 0-19.5-9.5t-3.5-30.5v-6l61-61l-158-158l-54 55q-18 18-43.5 18t-43.5-18q0-1-7-11l-122-122q-18-18-18-43.5t18-43.5t43.5-18t43.5 18l88 88l55-56q18-18 43.5-18t43.5 18q5 5 10 16l189 188l60-60h7q14 0 21 1t13 6t6 16v185q0 19-13.5 32.5t-33.5 13.5z"/>`),
		g.Group(children),
	)
}