package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Knob(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M128 896L0 704L768 0l192 64l64 192l-704 768zm503-448l265-265l-55-55l-265 265zM0 512q0-104 40.5-199t109-163.5T313 40.5T512 0q80 0 157 25L10 613Q0 562 0 512zm1024 0q0 104-40.5 199t-109 163.5t-163.5 109t-199 40.5q-51 0-101-10l588-659q25 77 25 157z"/>`),
		g.Group(children),
	)
}