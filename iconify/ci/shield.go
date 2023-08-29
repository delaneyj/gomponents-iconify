package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shield(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12.926 20.631C15.032 19.678 20 16.733 20 10.165V6.197c0-1.118 0-1.678-.218-2.105a2.001 2.001 0 0 0-.875-.874C18.48 3 17.92 3 16.8 3H7.2c-1.12 0-1.68 0-2.108.218a1.999 1.999 0 0 0-.874.874C4 4.52 4 5.08 4 6.2v3.965c0 6.568 4.968 9.513 7.074 10.466c.223.102.335.152.588.195c.16.028.518.028.677 0c.252-.043.363-.093.585-.194h.002Z"/>`),
		g.Group(children),
	)
}