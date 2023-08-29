package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CollectionText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M43 88v299h298v42H43q-18 0-30.5-12.5T0 387V88h43zM384 3q18 0 30.5 12.5T427 45v256q0 18-12.5 30.5T384 344H128q-18 0-30.5-12.5T85 301V45q0-17 12.5-29.5T128 3h256zm-21 192v-43H149v43h214zm-86 85v-43H149v43h128zm86-171V67H149v42h214z"/>`),
		g.Group(children),
	)
}