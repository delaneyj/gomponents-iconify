package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CollectionMusic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M384 3q18 0 30.5 12.5T427 45v256q0 18-12.5 30.5T384 344H128q-18 0-30.5-12.5T85 301V45q0-17 12.5-29.5T128 3h256zm-43 106V67h-85v117q-14-11-32-11q-22 0-37.5 16T171 227t15.5 37.5T224 280t37.5-15.5T277 227V109h64zM43 88v299h298v42H43q-18 0-30.5-12.5T0 387V88h43z"/>`),
		g.Group(children),
	)
}