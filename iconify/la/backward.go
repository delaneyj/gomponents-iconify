package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Backward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m17 5l-1.594 1.188l-12 9L2.344 16l1.062.813l12 9L17 27v-7.375l8.406 6.188L27 26.968V5.03l-1.594 1.157L17 12.375zm8 3.969V23.03l-8.406-6.187L15.469 16l1.125-.844zM15 9v14l-9.313-7z"/>`),
		g.Group(children),
	)
}