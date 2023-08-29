package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Blip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m2 408l33-34v-53l108 108H90l-34 33zm27-260v83q84 0 145 59q59 59 59 145h83q0-118-84.5-202.5T29 148zM29 2v82q144 0 247 103.5T379 435h83q0-88-34.5-168T335 129T197 36.5T29 2z"/>`),
		g.Group(children),
	)
}