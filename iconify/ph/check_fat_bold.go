package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckFatBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m246.15 65.46l-.07-.07l-23.93-23.54a20 20 0 0 0-28.22-.06l-90 88.62L70 97.76a20 20 0 0 0-28.19.1l-24 24a20 20 0 0 0 0 28.26l71.62 72a20 20 0 0 0 28.29 0L246.15 93.74a20 20 0 0 0 0-28.28ZM103.61 202.33L37.64 136l18.41-18.4L90 150.24a20 20 0 0 0 28.12 0L208 61.61l18.32 18Z"/>`),
		g.Group(children),
	)
}