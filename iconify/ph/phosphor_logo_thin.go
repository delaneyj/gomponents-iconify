package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhosphorLogoThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M144 28H64a4 4 0 0 0-4 4v128a76.08 76.08 0 0 0 76 76a4 4 0 0 0 4-4v-68h4a68 68 0 0 0 0-136ZM68 47.27L129.16 156H68Zm64 97.46L70.84 36H132ZM68.13 164H132v63.88A68.1 68.1 0 0 1 68.13 164Zm75.87-8h-4V36h4a60 60 0 0 1 0 120Z"/>`),
		g.Group(children),
	)
}