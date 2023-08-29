package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Onelist(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.14 34.07a2.64 2.64 0 1 1 2.64-2.64a2.64 2.64 0 0 1-2.64 2.64Zm0-7.43A2.65 2.65 0 1 1 9.78 24a2.65 2.65 0 0 1-2.64 2.65Zm0-7.43a2.64 2.64 0 1 1 2.64-2.64a2.64 2.64 0 0 1-2.64 2.64Zm5.61 12.17H43.5m-30.74-7.43h19.57m-19.58-7.43h9.01"/>`),
		g.Group(children),
	)
}