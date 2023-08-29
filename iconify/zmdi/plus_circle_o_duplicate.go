package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlusCircleODuplicate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M341 107v64h64v42h-64v64h-42v-64h-64v-42h64v-64h42zM43 192q0 44 23.5 80.5T128 327v46q-56-20-92-69.5T0 192T36 80.5T128 11v46q-38 18-61.5 54.5T43 192zM320 0q79 0 135.5 56.5T512 192t-56.5 135.5T320 384t-135.5-56.5T128 192t56.5-135.5T320 0zm0 341q62 0 105.5-43.5T469 192T425.5 86.5T320 43T214.5 86.5T171 192t43.5 105.5T320 341z"/>`),
		g.Group(children),
	)
}