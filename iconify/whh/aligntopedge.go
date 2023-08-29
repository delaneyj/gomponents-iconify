package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Aligntopedge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.12 128h-896q-26 0-45-18.5t-19-45T18.62 19t45.5-19h896q26 0 45 19t19 45.5t-18.5 45t-45.5 18.5zm-832 128h256q27 0 45.5 18.5t18.5 45.5v640q0 26-18.5 45t-45.5 19h-256q-26 0-45-19t-19-45V320q0-27 19-45.5t45-18.5zm512 0h256q27 0 45.5 18.5t18.5 45.5v384q0 26-18.5 45t-45.5 19h-256q-26 0-45-19t-19-45V320q0-27 19-45.5t45-18.5zm64 352q0 13 9.5 22.5t22.5 9.5h64q13 0 22.5-9.5t9.5-22.5V416q0-13-9.5-22.5t-22.5-9.5h-64q-13 0-22.5 9.5t-9.5 22.5v192z"/>`),
		g.Group(children),
	)
}