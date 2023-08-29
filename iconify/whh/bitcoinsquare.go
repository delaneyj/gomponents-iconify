package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bitcoinsquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-64-640q0-80-56-136t-136-56q0-27-18.5-45.5t-45-18.5t-45.5 18.5t-19 45.5h-64q0-27-18.5-45.5t-45-18.5t-45.5 18.5t-19 45.5h-64q-26 0-45 18.5t-19 45.5t19 45.5t45 18.5h64v384h-64q-26 0-45 19t-19 45.5t19 45t45 18.5h64q0 27 19 45.5t45.5 18.5t45-18.5t18.5-45.5h64q0 27 19 45.5t45.5 18.5t45-18.5t18.5-45.5q80 0 136-56t56-136q0-73-50-128q50-55 50-128zm-192 320h-192V576h192q27 0 45.5 19t18.5 45.5t-18.5 45t-45.5 18.5zm0-256h-192V320h192q27 0 45.5 18.5t18.5 45.5t-18.5 45.5t-45.5 18.5z"/>`),
		g.Group(children),
	)
}