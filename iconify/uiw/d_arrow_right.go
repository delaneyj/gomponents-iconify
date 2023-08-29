package uiw

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DArrowRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m2.542 2.154l7.254 7.26c.136.14.204.302.204.483a.73.73 0 0 1-.204.5l-7.575 7.398c-.383.317-.724.317-1.022 0c-.299-.317-.299-.643 0-.98l7.08-6.918l-6.754-6.763c-.237-.343-.215-.654.066-.935c.281-.28.598-.295.951-.045Zm9 0l7.254 7.26c.136.14.204.302.204.483a.73.73 0 0 1-.204.5l-7.575 7.398c-.383.317-.724.317-1.022 0c-.299-.317-.299-.643 0-.98l7.08-6.918l-6.754-6.763c-.237-.343-.215-.654.066-.935c.281-.28.598-.295.951-.045Z"/>`),
		g.Group(children),
	)
}