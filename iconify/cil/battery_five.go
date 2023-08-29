package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M80 160h32v200H80zm64 0h32v200h-32zm64 0h32v200h-32zm64 0h32v200h-32zm64 0h32v200h-32z"/><path fill="currentColor" d="M432 176v-56.41A23.825 23.825 0 0 0 408 96H39.986a23.825 23.825 0 0 0-24 23.59v280.82a23.825 23.825 0 0 0 24 23.59H408a23.825 23.825 0 0 0 24-23.59V344h64V176Zm32 136h-64v80H47.986V128H400v80h64Z"/>`),
		g.Group(children),
	)
}