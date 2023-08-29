package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockUnlocked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M384 200v-56a128 128 0 0 0-217.582-91.43l22.4 22.855A96 96 0 0 1 352 144v56H88v128c0 92.636 75.364 168 168 168s168-75.364 168-168V200Zm8 128c0 74.99-61.009 136-136 136s-136-61.01-136-136v-96h272Z"/>`),
		g.Group(children),
	)
}