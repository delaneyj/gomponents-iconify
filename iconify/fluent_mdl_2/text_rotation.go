package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextRotation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1332 640l-48 128h-103l240-640h102l240 640h-103l-48-128h-280zm140-375l-92 247h184l-92-247zm227 1242l90 90l-317 317l-317-317l90-90l163 163V896h128v774l163-163zM640 1670l163-163l90 90l-317 317l-317-317l90-90l163 163V128h128v1542z"/>`),
		g.Group(children),
	)
}