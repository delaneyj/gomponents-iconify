package flat_ui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Imac(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="#BDC3C7" fill-rule="evenodd" d="M95 79H62v12h4.5a1.5 1.5 0 1 1 0 3h-33a1.5 1.5 0 1 1 0-3H38V79H5a5 5 0 0 1-5-5v-9h100v9a5 5 0 0 1-5 5z" clip-rule="evenodd"/><path fill="#95A5A5" fill-rule="evenodd" d="M38 79h24v2H38v-2z" clip-rule="evenodd"/><path fill="#2C3E50" fill-rule="evenodd" d="M5 0h90a5 5 0 0 1 5 5v60H0V5a5 5 0 0 1 5-5z" clip-rule="evenodd"/><path fill="#35495E" fill-rule="evenodd" d="M5 5h90v55H5V5z" clip-rule="evenodd"/><circle cx="50" cy="72" r="2" fill="#35495E" fill-rule="evenodd" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}