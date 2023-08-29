package websymbol

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func List(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1000 1000"),
		g.Raw(`<path fill="currentColor" d="M1000 118v213H331V118h669zm0 334v213H331V452h669zM239 224q0 31-22.5 53.5T163 300t-53.5-22.5T87 224t22.5-53.5T163 148t53.5 22.5T239 224zm761 563v213H331V787h669zM239 559q0 32-22.5 54T163 635t-53.5-22T87 559t22.5-54t53.5-22t53.5 22t22.5 54zm0 335q0 31-22.5 53.5T163 970t-53.5-22.5T87 894q0-32 22.5-54.5T163 817t53.5 22.5T239 894z"/>`),
		g.Group(children),
	)
}