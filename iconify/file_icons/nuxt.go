package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nuxt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M200.662 81.35L0 430.65h130.774l139.945-239.803zm134.256 40.313l-39.023 69.167l138.703 239.82H512zm-51.596 91.052L155.924 430.651h253.485z"/>`),
		g.Group(children),
	)
}