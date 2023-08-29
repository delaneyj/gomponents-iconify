package bpmn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataStore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="70" d="M355.105 530.563v938.874c42.996 208.638 1246.884 208.638 1289.88 0V530.563c-42.996-208.638-1246.884-208.638-1289.88 0c42.996 208.639 1246.884 208.639 1289.88 0M355.105 708.61c42.996 208.639 1246.884 208.639 1289.88 0M355.105 886.657c42.996 208.639 1246.884 208.639 1289.88 0"/>`),
		g.Group(children),
	)
}