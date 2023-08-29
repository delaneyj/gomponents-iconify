package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayOrPauseButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M53.213 10.788c-11.715-11.717-30.711-11.717-42.426 0c-11.716 11.715-11.716 30.711 0 42.426c11.715 11.715 30.711 11.715 42.426 0c11.716-11.715 11.716-30.711 0-42.426M13 47.999v-32l19 16l-19 16m28.5 0h-5.7v-32h5.7v32m9.5 0h-5.7v-32H51v32"/>`),
		g.Group(children),
	)
}