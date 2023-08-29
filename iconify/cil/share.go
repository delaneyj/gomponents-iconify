package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Share(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M68.146 496H16V235.333a114.169 114.169 0 0 1 12.025-51.309A136.27 136.27 0 0 1 152 104h136.557V16h42.361l163.709 163.71L330.337 344h-41.78v-88h-88.812a85.4 85.4 0 0 0-81.993 62.244ZM152 136a104.217 104.217 0 0 0-94.923 61.443l-.292.614A82.454 82.454 0 0 0 48 235.333v213.81l38.93-139.5A117.5 117.5 0 0 1 199.745 224h120.812v84.525L449.373 179.71L320.557 50.894V136Z"/>`),
		g.Group(children),
	)
}