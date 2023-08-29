package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Monitor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.193 768h-384v128h64q53 0 90.5 37.5t37.5 90.5h-512q0-53 37.5-90.5t90.5-37.5h64V768h-384q-27 0-45.5-18.5T.193 704V64q0-26 18.5-45t45.5-19h896q26 0 45 19t19 45v640q0 27-18.5 45.5t-45.5 18.5zm-64-640h-768v512h768V128zm-704 448V191h512z"/>`),
		g.Group(children),
	)
}