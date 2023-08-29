package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lego(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M0 704V256q0-27 18.5-45.5T64 192h896q26 0 45 18.5t19 45.5v448H0zm640-576V64q0-26 18.5-45T704 0h128q26 0 45 19t19 45v64H640zm-512 0V64q0-26 18.5-45T192 0h128q27 0 45.5 19T384 64v64H128z"/>`),
		g.Group(children),
	)
}