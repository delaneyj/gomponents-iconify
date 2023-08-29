package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Label(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m19.293 9.951l-2.333-2.8c-.353-.423-.53-.635-.746-.787a2.001 2.001 0 0 0-.632-.295C15.327 6 15.052 6 14.502 6H7.2c-1.12 0-1.68 0-2.108.218a1.999 1.999 0 0 0-.874.874C4 7.52 4 8.08 4 9.2v5.6c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218H14.5c.551 0 .826 0 1.081-.069c.226-.06.44-.16.632-.296c.216-.152.393-.363.746-.786l2.333-2.8c.608-.729.91-1.093 1.027-1.5c.102-.359.102-.74 0-1.098c-.116-.407-.42-.77-1.027-1.5Z"/>`),
		g.Group(children),
	)
}