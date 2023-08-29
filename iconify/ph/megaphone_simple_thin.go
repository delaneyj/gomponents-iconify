package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MegaphoneSimpleThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m219.41 90.5l-176-54A12 12 0 0 0 28 48v144a12 12 0 0 0 12 12a12.41 12.41 0 0 0 3.45-.5L132 176.32V192a12 12 0 0 0 12 12h32a12 12 0 0 0 12-12v-32.86l31.39-9.63A12.06 12.06 0 0 0 228 138v-36a12 12 0 0 0-8.59-11.5ZM41.12 195.84A4 4 0 0 1 36 192V48a4 4 0 0 1 1.6-3.2A4 4 0 0 1 40 44a3.89 3.89 0 0 1 1.07.15L132 72.05V168ZM180 192a4 4 0 0 1-4 4h-32a4 4 0 0 1-4-4v-18.13l40-12.27Zm40-54a4 4 0 0 1-2.88 3.84h-.05L140 165.5v-91l77.12 23.66A4 4 0 0 1 220 102Z"/>`),
		g.Group(children),
	)
}