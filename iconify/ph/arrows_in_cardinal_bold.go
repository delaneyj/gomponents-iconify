package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsInCardinalBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M87.51 72.49a12 12 0 0 1 17-17L116 67V24a12 12 0 0 1 24 0v43l11.51-11.52a12 12 0 0 1 17 17l-32 32a12 12 0 0 1-17 0Zm49 79a12 12 0 0 0-17 0l-32 32a12 12 0 0 0 17 17L116 189v43a12 12 0 0 0 24 0v-43l11.51 11.52a12 12 0 0 0 17-17ZM232 116h-43l11.52-11.51a12 12 0 0 0-17-17l-32 32a12 12 0 0 0 0 17l32 32a12 12 0 0 0 17-17L189 140h43a12 12 0 0 0 0-24Zm-127.51 3.51l-32-32a12 12 0 0 0-17 17L67 116H24a12 12 0 0 0 0 24h43l-11.49 11.51a12 12 0 0 0 17 17l32-32a12 12 0 0 0-.02-17Z"/>`),
		g.Group(children),
	)
}