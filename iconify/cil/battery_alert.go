package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryAlert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M432 176v-56.41A23.825 23.825 0 0 0 408 96H240v32h160v80h64v104h-64v80H47.986V128H112V96H39.986a23.825 23.825 0 0 0-24 23.59v280.82a23.825 23.825 0 0 0 24 23.59H408a23.825 23.825 0 0 0 24-23.59V344h64V176Z"/><path fill="currentColor" d="M192 64v-8h-32v200h32V64zm0 240v-16h-32v32h32v-16z"/>`),
		g.Group(children),
	)
}