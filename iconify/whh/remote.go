package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Remote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-640-695q-8-9-20-9t-21 9l-143 161q-8 9-8 22t8 22l143 161q9 9 21 9t20-9V329zm439-114l-161-143q-9-8-22-8t-22 8l-161 143q-9 8-9 20.5t9 20.5h366q9-8 9-20.5t-9-20.5zm-375 297q0 80 56 136t136 56t136-56.5t56-136t-56-135.5t-136-56t-136 56t-56 136zm375 255h-366q-9 8-9 20.5t9 20.5l161 143q9 8 22 8t22-8l161-143q9-8 9-20.5t-9-20.5zm257-277l-143-161q-9-9-21-9t-20 9v366q8 9 20 9t21-9l143-161q8-9 8-22t-8-22z"/>`),
		g.Group(children),
	)
}