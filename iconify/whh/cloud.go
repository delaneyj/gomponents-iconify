package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cloud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M768 768H224q-93 0-158.5-65.5T0 544q0-57 27-106t73-80q34 69 96.5 116.5T336 537q-66-38-105-104t-39-145q0-119 84.5-203.5T480 0q88 0 159.5 48T744 174q-68 22-118 73.5T556 369q35-52 91-82.5T768 256q106 0 181 75t75 181t-75 181t-181 75z"/>`),
		g.Group(children),
	)
}