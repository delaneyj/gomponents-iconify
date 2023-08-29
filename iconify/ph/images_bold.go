package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImagesBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 36H80a20 20 0 0 0-20 20v12H40a20 20 0 0 0-20 20v112a20 20 0 0 0 20 20h144a20 20 0 0 0 20-20v-20h12a20 20 0 0 0 20-20V56a20 20 0 0 0-20-20ZM84 60h128v45.09l-3.23-3.23a20 20 0 0 0-28.28 0L165.31 117l-35.17-35.14a20 20 0 0 0-28.28 0L84 99.72Zm96 136H44V92h16v68a20 20 0 0 0 20 20h100Zm-96-40v-22.34l32-32l40.83 40.83a12 12 0 0 0 17 0l20.83-20.83L212 139v17Z"/>`),
		g.Group(children),
	)
}