package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tickertape(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="5.606" height="3.441" x="5.5" y="22.279" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx=".781"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m18.283 27.911l-4.516-7.822H22.8Zm8.873-7.947l-4.516 7.822h9.033Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 5.5a2 2 0 0 0-2 2v9.295h27.502a1.393 1.393 0 0 1 1.396 1.396V29.81a1.393 1.393 0 0 1-1.396 1.396H5.5V40.5a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}