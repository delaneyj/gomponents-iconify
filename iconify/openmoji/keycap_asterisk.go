package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeycapAsterisk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#d0cfce" d="M11.875 12.166h48V60h-48z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12.125 11.916h48v48h-48zm24 13v22M26.923 29.89l18.404 12.053m-18.415-.017l18.426-12.02"/>`),
		g.Group(children),
	)
}