package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayOrPauseButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#40c0e7" d="M67.68 95.55h18.03v-63.1H67.68v63.1zm36.05-63.1v63.09h18.03V32.45h-18.03zm-97.49-.28v63.66L56.26 64L6.24 32.17z"/>`),
		g.Group(children),
	)
}