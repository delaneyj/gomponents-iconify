package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Blemonitor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.14 5.5h31.72a2.65 2.65 0 0 1 2.64 2.64v31.72a2.65 2.65 0 0 1-2.64 2.64H8.14a2.65 2.65 0 0 1-2.64-2.64V8.14A2.65 2.65 0 0 1 8.14 5.5Z"/><circle cx="17.45" cy="25.4" r="1.3" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.45 19.77a5.59 5.59 0 0 1 5.63 5.63m-5.63-10.62A10.71 10.71 0 0 1 28.07 25.4"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.45 9.58A15.74 15.74 0 0 1 33.16 25.4M17.5 30.5v8.2h4.1m2 0h4.1m-4.1-8.2h4.1m-4.1 4.1h2.7m-2.7-4.1v8.2m6.2 0v-8.2l4.1 8.2l4.1-8.2v8.2m-24.6-4.1a2 2 0 0 1 0 4H10v-8.2h3.4a2 2 0 0 1 2 2a2.19 2.19 0 0 1-2 2.2Zm0 0H10"/>`),
		g.Group(children),
	)
}