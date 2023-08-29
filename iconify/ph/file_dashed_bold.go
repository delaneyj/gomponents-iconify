package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileDashedBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M84 224a12 12 0 0 1-12 12H56a20 20 0 0 1-20-20v-32a12 12 0 0 1 24 0v28h12a12 12 0 0 1 12 12ZM220 88v48a12 12 0 0 1-24 0v-32h-48a12 12 0 0 1-12-12V44h-16a12 12 0 0 1 0-24h32a12 12 0 0 1 8.49 3.51l56 56A12 12 0 0 1 220 88Zm-60-8h23l-23-23ZM80 20H56a20 20 0 0 0-20 20v24a12 12 0 0 0 24 0V44h20a12 12 0 0 0 0-24Zm128 144a12 12 0 0 0-12 12v36h-4a12 12 0 0 0 0 24h8a20 20 0 0 0 20-20v-40a12 12 0 0 0-12-12Zm-160-8a12 12 0 0 0 12-12v-40a12 12 0 0 0-24 0v40a12 12 0 0 0 12 12Zm104 56h-40a12 12 0 0 0 0 24h40a12 12 0 0 0 0-24Z"/>`),
		g.Group(children),
	)
}