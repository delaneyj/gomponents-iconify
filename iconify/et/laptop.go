package et

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Laptop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M3.5 26a.5.5 0 0 0 .5-.5V2c0-.598.402-1 1-1h32c.57 0 1 .468 1 1.089V25.5a.5.5 0 0 0 1 0V2.089C39 .917 38.122 0 37 0H5C3.841 0 3 .841 3 2v23.5a.5.5 0 0 0 .5.5z"/><path d="M6.5 3a.5.5 0 0 0-.5.5v22a.5.5 0 0 0 .5.5h29a.5.5 0 0 0 .5-.5v-22a.5.5 0 0 0-.5-.5h-29zM35 25H7V4h28v21zm6.5 3H.5a.5.5 0 0 0-.5.5v1.175C0 30.873 1.408 32 2.5 32h37c1.145 0 2.5-1.355 2.5-2.5v-1a.5.5 0 0 0-.5-.5zm-.5 1.5c0 .589-.911 1.5-1.5 1.5h-37c-.631 0-1.5-.768-1.5-1.325V29h40v.5z"/></g>`),
		g.Group(children),
	)
}