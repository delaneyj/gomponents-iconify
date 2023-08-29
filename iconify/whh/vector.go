package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vector(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.356 1024h-64q-27 0-45.5-19t-18.5-45v-64q0-27 18.5-45.5t45.5-18.5q0-111-73.5-192.5t-182.5-93.5v30q0 26-18.5 45t-45.5 19h-128q-27 0-45.5-19t-18.5-45v-34q-135-12-227.5-112.5T64.356 192q-26 0-45-19t-19-45V64q0-27 19-45.5t45-18.5h64q27 0 45.5 18.5t18.5 45.5v64q0 26-19 45t-45 19q0 111 73.5 192.5t182.5 93.5v-30q0-27 18.5-45.5t45.5-18.5h128q27 0 45.5 18.5t18.5 45.5v34q135 12 227.5 112.5t92.5 237.5q26 0 45 18.5t19 45.5v64q0 27-19 45.5t-45 18.5zm-448.5-576q-26.5 0-45 19t-18.5 45.5t18.5 45t45 18.5t45.5-18.5t19-45t-19-45.5t-45.5-19z"/>`),
		g.Group(children),
	)
}