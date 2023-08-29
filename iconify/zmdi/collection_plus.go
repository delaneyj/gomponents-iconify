package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CollectionPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M43 88v299h298v42H43q-18 0-30.5-12.5T0 387V88h43zM384 3q18 0 30.5 12.5T427 45v256q0 18-12.5 30.5T384 344H128q-18 0-30.5-12.5T85 301V45q0-17 12.5-29.5T128 3h256zm-21 192v-43h-86V67h-42v85h-86v43h86v85h42v-85h86z"/>`),
		g.Group(children),
	)
}