package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandPeanut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 16.25l-.816-.36l-.462-.196c-1.444-.592-2-.593-3.447 0l-.462.195l-.817.359a4.5 4.5 0 1 1 0-8.49v0l1.054.462l.434.178c1.292.507 1.863.48 3.237-.082l.462-.195l.817-.359a4.5 4.5 0 1 1 0 8.49"/>`),
		g.Group(children),
	)
}