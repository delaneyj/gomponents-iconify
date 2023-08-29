package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerMouse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 192h176V0h-16C71.6 0 0 71.6 0 160v32zm0 32v128c0 88.4 71.6 160 160 160h64c88.4 0 160-71.6 160-160V224H0zm384-32v-32C384 71.6 312.4 0 224 0h-16v192h176z"/>`),
		g.Group(children),
	)
}