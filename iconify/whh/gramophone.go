package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gramophone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M992 192h-32v416q0 1-.5 3t-.5 2.5v4.5l-56 253q-2 7-7 12v45q0 13-9.5 22.5T864 960h-64q-13 0-22.5-9.5T768 928v-64q0-15 82-85l46-203V192h-32q-13 0-22.5-9.5T832 160V32q0-13 9.5-22.5T864 0h128q13 0 22.5 9.5T1024 32v128q0 13-9.5 22.5T992 192zM416 960q-113 0-209-55.5T55.5 753T0 544t55.5-209T207 183.5T416 128t209 55.5T776.5 335T832 544t-55.5 209T625 904.5T416 960zm0-576q-66 0-113 47t-47 113t47 113t113 47t113-47t47-113t-47-113t-113-47zm0 192q-13 0-22.5-9.5T384 544t9.5-22.5T416 512t22.5 9.5T448 544t-9.5 22.5T416 576z"/>`),
		g.Group(children),
	)
}