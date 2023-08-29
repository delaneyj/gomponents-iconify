package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trashempty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.743 192h-896q-26 0-45-18.5t-19-45t19-45.5t45-19h320q0-26 18.5-45t45.5-19h128q27 0 45.5 19t18.5 45h320q26 0 45 19t19 45.5t-19 45t-45 18.5zm-832 64q27 0 45.5 19t18.5 45v512q0 27 19 45.5t45 18.5h512q26 0 45-18.5t19-45.5V320q0-26 18.5-45t45-19t45.5 19t19 45v576q0 53-37.5 90.5t-90.5 37.5h-640q-53 0-90.5-37.5t-37.5-90.5V320q0-26 18.5-45t45.5-19z"/>`),
		g.Group(children),
	)
}