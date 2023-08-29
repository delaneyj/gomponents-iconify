package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadioThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M100 176a4 4 0 0 1-4 4H64a4 4 0 0 1 0-8h32a4 4 0 0 1 4 4Zm-4-36H64a4 4 0 0 0 0 8h32a4 4 0 0 0 0-8Zm132-52v112a12 12 0 0 1-12 12H40a12 12 0 0 1-12-12V80a4 4 0 0 1 2.85-3.81l160-48a4 4 0 0 1 2.3 7.66L59.25 76H216a12 12 0 0 1 12 12Zm-8 0a4 4 0 0 0-4-4H36v116a4 4 0 0 0 4 4h176a4 4 0 0 0 4-4Zm-24 56a36 36 0 1 1-36-36a36 36 0 0 1 36 36Zm-8 0a28 28 0 1 0-28 28a28 28 0 0 0 28-28Zm-92-36H64a4 4 0 0 0 0 8h32a4 4 0 0 0 0-8Z"/>`),
		g.Group(children),
	)
}