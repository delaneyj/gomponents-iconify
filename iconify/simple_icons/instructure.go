package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Instructure(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11.996 0l-5.11 2.878L12 5.76l5.115-2.878ZM6.032 3.36L.918 6.237L6.036 9.12l5.115-2.879Zm11.929 0l-5.112 2.878l5.115 2.882l5.118-2.879zM12 11.52L.918 17.76L12 24l11.082-6.241Z"/>`),
		g.Group(children),
	)
}