package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mailinglists(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M831.405 1024h-767q-27 0-45.5-18.5T.405 960V65q0-27 18.5-45.5T64.405 1h448v255h-352q-13 0-22.5 9t-9.5 23v64q0 13 9.5 22.5t22.5 9.5l735 1v575q0 27-18.5 45.5t-45.5 18.5zm-32-480q0-14-9-23t-22-9h-608q-13 0-22.5 9t-9.5 23v64q0 13 9.5 22.5t22.5 9.5h608q13 0 22-9.5t9-22.5v-64zm0 255q0-13-9-22.5t-22-9.5h-608q-13 0-22.5 9.5t-9.5 22.5v64q0 14 9.5 23t22.5 9h608q13 0 22-9t9-23v-64zm-223-799q26 0 44 18l256 257q19 19 19 46h-319V0z"/>`),
		g.Group(children),
	)
}