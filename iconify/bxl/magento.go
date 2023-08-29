package bxl

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Magento(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 22.019l-3.717-2.146V9.863l2.479-1.43v10.01l1.238.753l1.238-.753V8.434l2.479 1.43v10.01L12 22.019zm8.666-15.014v10.009l-2.475 1.43V8.434L12 4.861L5.807 8.434v10.01l-2.473-1.43V7.005L12 2l8.666 5.005z"/>`),
		g.Group(children),
	)
}