package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Riscv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.95.051h16.947v6.21L13.473 21.004l-.524.734l-5.789-6.842c4.103-.74 6.21-3.896 6.21-7.37C13.37 4.05 11.263.472 6.95.05zm-5.475 13.37l8.74 10.528H0V3.419h5.474c2.945 0 4.422 1.999 4.422 4.107c0 2.107-1.477 4.21-4.422 4.21H1.475v1.685zm14.07 10.528H24V12.157l-7.685 10.738l-.77 1.054z"/>`),
		g.Group(children),
	)
}