package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Turbolinux(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m9.419 6.222l.547 1.23h-3.35L6 6.223h3.419Zm3.692 5.949L7.094 0l7.042 4.17l.41 1.984h3.351l-.752 2.051h-2.188l1.778 8.274l-4.171-2.052L14.684 24L8.187 10.803l4.923 1.368Z"/>`),
		g.Group(children),
	)
}