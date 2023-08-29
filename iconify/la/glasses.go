package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Glasses(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m5.063 6l-.25.656L1.313 16H0v2h1v2c0 3.297 2.703 6 6 6h2c3.297 0 6-2.703 6-6v-3c0-.566.434-1 1-1c.566 0 1 .434 1 1v3c0 3.297 2.703 6 6 6h2c3.297 0 6-2.703 6-6v-2h1v-2h-1.313l-3.5-9.344l-.25-.656H25l-.281.375l-2.5 3l1.562 1.25l1.907-2.313L28.563 16h-9.75c-.418-1.156-1.52-2-2.813-2c-1.293 0-2.395.844-2.813 2h-9.75l2.876-7.688l1.906 2.313l1.562-1.25l-2.5-3L7 6zM3 18h10v2c0 2.203-1.797 4-4 4H7c-2.203 0-4-1.797-4-4zm16 0h10v2c0 2.203-1.797 4-4 4h-2c-2.203 0-4-1.797-4-4z"/>`),
		g.Group(children),
	)
}