package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BasketThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M132 120v56a4 4 0 0 1-8 0v-56a4 4 0 0 1 8 0Zm103.89-22.41l-13.89 104A12 12 0 0 1 210.13 212H45.87A12 12 0 0 1 34 201.59l-13.86-104A12 12 0 0 1 32 84h38.18L125 21.37a4 4 0 0 1 6 0L185.82 84H224a12 12 0 0 1 11.89 13.59ZM80.81 84h94.38L128 30.07ZM227 93.36a3.94 3.94 0 0 0-3-1.36H32a4 4 0 0 0-4 4.52l13.86 104a4 4 0 0 0 4 3.47h164.27a4 4 0 0 0 4-3.47L228 96.53a3.94 3.94 0 0 0-1-3.17ZM181.2 116a4 4 0 0 0-4.38 3.58l-5.6 56a4 4 0 0 0 3.58 4.42h.41a4 4 0 0 0 4-3.6l5.6-56a4 4 0 0 0-3.61-4.4Zm-106.4 0a4 4 0 0 0-3.58 4.38l5.6 56a4 4 0 0 0 4 3.6h.41a4 4 0 0 0 3.58-4.38l-5.6-56a4 4 0 0 0-4.41-3.6Z"/>`),
		g.Group(children),
	)
}