package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignInDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="m136 128l-40 40V88Z" opacity=".2"/><path d="m141.66 122.34l-40-40A8 8 0 0 0 88 88v32H24a8 8 0 0 0 0 16h64v32a8 8 0 0 0 13.66 5.66l40-40a8 8 0 0 0 0-11.32ZM104 148.69v-41.38L124.69 128ZM208 48v160a16 16 0 0 1-16 16h-56a8 8 0 0 1 0-16h56V48h-56a8 8 0 0 1 0-16h56a16 16 0 0 1 16 16Z"/></g>`),
		g.Group(children),
	)
}