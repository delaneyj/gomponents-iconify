package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DummySquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 9.2v5.6c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h5.606c1.118 0 1.677 0 2.104-.218c.377-.192.683-.498.875-.874c.218-.428.218-.987.218-2.104V9.197c0-1.118 0-1.678-.218-2.105a2.001 2.001 0 0 0-.875-.874C16.48 6 15.92 6 14.8 6H9.2c-1.12 0-1.68 0-2.108.218a1.999 1.999 0 0 0-.874.874C6 7.52 6 8.08 6 9.2Z"/>`),
		g.Group(children),
	)
}