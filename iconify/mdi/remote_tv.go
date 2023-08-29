package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RemoteTv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 2c-1.11 0-2 .89-2 2v16c0 1.11.89 2 2 2h6c1.11 0 2-.89 2-2V4c0-1.11-.89-2-2-2h-2v2h-2V2H9m2 4h2v2h2v2h-2v2h-2v-2H9V8h2V6m-2 8h2v2H9v-2m4 0h2v2h-2v-2m-4 4h2v2H9v-2m4 0h2v2h-2v-2Z"/>`),
		g.Group(children),
	)
}