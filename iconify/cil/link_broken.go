package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M288 496h32V368h128v-32H288v160zm-64-320V16h-32v128H64v32h160zM78.708 434.573l-1.279-1.273a100.478 100.478 0 0 1 0-142.1l75.2-75.2h-45.257L54.8 268.57a132.478 132.478 0 0 0 0 187.35l1.278 1.278a132.628 132.628 0 0 0 187.352 0l4.57-4.57v-45.255l-27.2 27.2a100.591 100.591 0 0 1-142.092 0ZM457.2 56.08l-1.278-1.28c-51.653-51.655-135.7-51.653-187.352 0L264 59.372v45.255l27.2-27.2a100.591 100.591 0 0 1 142.095 0l1.279 1.278a100.478 100.478 0 0 1 0 142.1l-75.2 75.2h45.253L457.2 243.43a132.478 132.478 0 0 0 0-187.35Z"/>`),
		g.Group(children),
	)
}