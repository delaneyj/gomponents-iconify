package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Key(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M832 384q0-80-56-136t-136-56t-136 56t-56 136q0 42 19 83q-41-19-83-19q-80 0-136 56t-56 136t56 136t136 56t136-56t56-136q0-42-19-83q41 19 83 19q80 0 136-56t56-136zm851 704q0 17-49 66t-66 49q-9 0-28.5-16t-36.5-33t-38.5-40t-24.5-26l-96 96l220 220q28 28 28 68q0 42-39 81t-81 39q-40 0-68-28L733 893q-176 131-365 131q-163 0-265.5-102.5T0 656q0-160 95-313T343 95T656 0q163 0 265.5 102.5T1024 368q0 189-131 365l355 355l96-96q-3-3-26-24.5t-40-38.5t-33-36.5t-16-28.5q0-17 49-66t66-49q13 0 23 10q6 6 46 44.5t82 79.5t86.5 86t73 78t28.5 41z"/>`),
		g.Group(children),
	)
}