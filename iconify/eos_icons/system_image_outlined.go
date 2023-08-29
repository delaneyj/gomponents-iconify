package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SystemImageOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 6v2H7V6Zm-2 4v2h6v-2Zm-3 4v2h5v-2Zm3 4v2h6v-2Zm-4 0v2h2v-2Zm8-4v2h2v-2Zm0-8v2h2V6Zm4 16H5V4h14Zm0-20H5a2 2 0 0 0-2 2v18a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V0Z"/>`),
		g.Group(children),
	)
}