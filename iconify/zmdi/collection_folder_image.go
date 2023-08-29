package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CollectionFolderImage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M43 88v299h384v42H43q-18 0-30.5-12.5T0 387V88h43zm426-43q18 0 30.5 12.5T512 88v213q0 18-12.5 30.5T469 344H128q-18 0-30.5-12.5T85 301l1-256q0-17 12-29.5T128 3h128l43 42h170zM149 280h299l-75-96l-53 64l-75-96z"/>`),
		g.Group(children),
	)
}