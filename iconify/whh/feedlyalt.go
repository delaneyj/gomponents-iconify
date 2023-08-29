package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Feedlyalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-198-814l-78-77q-5-5-12.5-5t-12.5 5l-334 332q-6 6-6 13.5t6 13.5l59 52h64l314-308q6-5 6-12.5t-6-13.5zm0 311l-78-78q-5-5-12.5-5t-12.5 5l-202 201q-6 5-6 12.5t6 13.5l55 66h64l186-189q6-6 6-13.5t-6-12.5zm0 304l-78-84q-5-6-12.5-6t-13.5 6l-78 84q-5 6-5 14.5t5 13.5l60 43h64l58-43q6-5 6-13.5t-6-14.5z"/>`),
		g.Group(children),
	)
}