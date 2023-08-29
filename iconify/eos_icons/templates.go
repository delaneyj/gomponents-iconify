package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Templates(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 2H5a2.005 2.005 0 0 0-2 2v16a2.005 2.005 0 0 0 2 2h14a2.005 2.005 0 0 0 2-2V4a2.005 2.005 0 0 0-2-2Zm0 11.152v5.696L14 16Zm-7 1.694L7 12h10ZM5 4h14v2H5Zm0 4h8v2H5Zm5 8l-5 2.848v-5.704Zm2 1.15L17 20H7Z"/>`),
		g.Group(children),
	)
}