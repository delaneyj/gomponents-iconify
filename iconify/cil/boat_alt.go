package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoatAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M410.866 181.063A32.1 32.1 0 0 0 380.793 160h-39.239l-16-96H221.112l16 96h-39.558l-16-96H77.112l16 96H48v88H16v200h427.727L496 282.466V248h-60.793ZM298.446 96l10.667 64h-39.559l-10.666-64Zm-144 0l10.667 64h-39.559l-10.666-64ZM80 192h300.793l20.363 56H80Zm383.222 88l-42.949 136H48V280Z"/><path fill="currentColor" d="M80 320h32v32H80zm64 0h32v32h-32zm64 0h32v32h-32zm64 0h32v32h-32zm64 0h32v32h-32z"/>`),
		g.Group(children),
	)
}