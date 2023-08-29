package devicon_plain

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vagrant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="currentColor" d="M122.51 23.9V12.82l-26.6 15.47v9.35L74.64 83.55l-10.65 7.33v33.61l22.81-13.15l35.71-87.44zM63.99 66.94L48.04 29.71V19.18l-.11-.05l-15.85 9.16v9.35l21.28 47.92l10.63-5.26V66.94z"/><path fill="currentColor" d="M106.56 3.51L79.97 19.08l-.02-.01v10.64L63.99 66.94v12.45l-10.63 6.17l-21.28-47.92v-9.37l15.97-9.18l-26.6-15.58l-15.96 9.31v11.41l35.91 87.21l22.59 13.05V91.73l10.65-6.17l-.14-.08l21.41-47.84v-9.35l.01-.02l26.59-15.45l-15.95-9.31z"/>`),
		g.Group(children),
	)
}