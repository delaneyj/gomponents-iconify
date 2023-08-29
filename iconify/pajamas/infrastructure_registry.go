package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InfrastructureRegistry(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m8 2.732l-2.945 1.7L8 6.135l2.945-1.701L8 2.733Zm4.445.834L8 1L3.555 3.566l-1.43-.825a.75.75 0 1 0-.75 1.298l1.429.826V10l4.446 2.567v1.683a.75.75 0 0 0 1.5 0v-1.683L13.196 10V4.865l1.43-.826a.75.75 0 0 0-.751-1.298l-1.43.825Zm-.749 2.165L8.75 7.433v3.402l2.946-1.701V5.731ZM4.304 9.134l2.946 1.7v-3.4L4.304 5.73v3.403Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}