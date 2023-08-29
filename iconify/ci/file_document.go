package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileDocument(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 17h6m-6-3h6M13 3H8.2c-1.12 0-1.68 0-2.108.218a1.999 1.999 0 0 0-.874.874C5 4.52 5 5.08 5 6.2v11.6c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h7.606c1.118 0 1.677 0 2.104-.218c.377-.192.683-.498.875-.874c.218-.428.218-.986.218-2.104V9.326L18.999 9M13 3c.286.003.466.014.638.055c.204.05.4.13.579.24c.202.124.375.297.72.643l3.126 3.125c.346.346.518.518.642.72c.11.18.19.374.24.578c.04.173.051.354.054.639M13 3v2.8c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h2.802"/>`),
		g.Group(children),
	)
}