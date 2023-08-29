package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Journey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M24 40C32.2843 40 39 33.2843 39 25C39 16.7157 32.2843 10 24 10C15.7157 10 9 16.7157 9 25C9 33.2843 15.7157 40 24 40Z"/><path d="M20.0002 11C21.8049 12.0083 23.5002 13.5 23.3581 15.4454C23.2445 17 21.915 17.3469 21.6369 18.4711C21.359 19.5953 22.9674 20.8217 20.2464 22.636C18.4324 23.8455 12.971 26.1447 9 24"/><path stroke-linejoin="round" d="M9.50011 24.5C6.50011 26.3883 2.06824 31.521 4.00056 35C6.5 39.5 16.0004 35.6906 27.0001 27C37.9999 18.3094 42.2291 5.60642 42.2291 5.60642L37.0005 7"/><path d="M26.0006 40C26.0006 40 26.5003 35 30 32C33.4997 29 38.0007 29 38.0007 29"/></g>`),
		g.Group(children),
	)
}