package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Exit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m1006.428 555l-163 200q-11 12-29.5 12.5t-32-7.5t-13.5-19V640h-128q-26 0-45-18.5t-19-45.5V448q0-27 19-45.5t45-18.5h128V284q0-12 13.5-20t32-7.5t29.5 12.5l163 199q18 18 18 43.5t-18 43.5zm-365.5-299q-26.5 0-45.5-19t-19-45.5t-18.5-45t-45.5-18.5h-320q-26 0-45 18.5t-19 45.5v640q0 26 19 45t45 19h320q27 0 45.5-19t18.5-45.5t19-45t45.5-18.5t45 18.5t18.5 45.5v64q0 53-37.5 90.5t-90.5 37.5h-448q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h448q53 0 90.5 37.5t37.5 90.5v64q0 26-18.5 45t-45 19z"/>`),
		g.Group(children),
	)
}