package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MedicalCase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.435 6.465h-1.43v-.9a2.5 2.5 0 0 0-2.5-2.5h-5a2.5 2.5 0 0 0-2.5 2.5v.9h-1.44a2.5 2.5 0 0 0-2.5 2.5v9.47a2.5 2.5 0 0 0 2.5 2.5h12.87a2.5 2.5 0 0 0 2.5-2.5v-9.47a2.5 2.5 0 0 0-2.5-2.5Zm-10.43-.9a1.5 1.5 0 0 1 1.5-1.5h5a1.5 1.5 0 0 1 1.5 1.5v.9h-8Zm11.93 12.87a1.5 1.5 0 0 1-1.5 1.5H5.565a1.5 1.5 0 0 1-1.5-1.5v-9.47a1.5 1.5 0 0 1 1.5-1.5h12.87a1.5 1.5 0 0 1 1.5 1.5Z"/><path fill="currentColor" d="M14.505 13.675a.5.5 0 0 1-.5.5h-1.5v1.5a.5.5 0 0 1-.5.5a.5.5 0 0 1-.5-.5v-1.5h-1.5a.5.5 0 0 1-.5-.5a.5.5 0 0 1 .5-.5h1.5v-1.5a.5.5 0 0 1 .5-.5a.508.508 0 0 1 .5.5v1.5h1.5a.508.508 0 0 1 .5.5Z"/>`),
		g.Group(children),
	)
}