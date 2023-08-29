package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gender(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1009 512L864 961q-14 26-40 44.5t-58 19.5q-28-1-56-20.5T672 961L527 512q-23-58-9.5-93t49.5-35h402q36 0 49.5 35t-9.5 93zm-246 513h6h-6zM576 192q0-80 56.5-136t136-56T904 56t56 136t-56 136t-135.5 56t-136-56T576 192zm-119 833H54q-35 0-48.5-35t8.5-93l176-525q-56-21-91.5-70T63 192q0-80 56.5-136t136-56t136 56T448 192q0 62-36.5 111.5T318 373l179 524q23 58 9 93t-49 35z"/>`),
		g.Group(children),
	)
}