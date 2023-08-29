package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Iot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 13v9H3v-9Zm18 0v2h-2v7h-2v-7h-2v-2Zm-11-2a7.537 7.537 0 0 1 3.96 1.149l1.447-1.45A9.522 9.522 0 0 0 12 9a9.363 9.363 0 0 0-5.333 1.68l1.449 1.453A7.36 7.36 0 0 1 12 11Z"/><path fill="currentColor" d="M12 7a11.494 11.494 0 0 1 6.834 2.27l1.427-1.43A13.48 13.48 0 0 0 12 5a13.333 13.333 0 0 0-8.186 2.822l1.426 1.43A11.343 11.343 0 0 1 12 7Z"/><path fill="currentColor" d="M12 3a15.471 15.471 0 0 1 9.687 3.41l1.427-1.429A17.43 17.43 0 0 0 .96 4.964l1.427 1.429A15.328 15.328 0 0 1 12 3Zm0 10a4.5 4.5 0 1 0 4.5 4.5A4.5 4.5 0 0 0 12 13Zm0 7a2.5 2.5 0 1 1 2.5-2.5A2.5 2.5 0 0 1 12 20Z"/>`),
		g.Group(children),
	)
}