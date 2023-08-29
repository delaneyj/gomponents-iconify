package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CollectionImage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M427 301q0 18-12.5 30.5T384 344H128q-18 0-30.5-12.5T85 301V45q0-17 12.5-29.5T128 3h256q18 0 30.5 12.5T427 45v256zm-235-85l-64 85h256l-85-106l-64 79zM0 88h43v299h298v42H43q-18 0-30.5-12.5T0 387V88z"/>`),
		g.Group(children),
	)
}