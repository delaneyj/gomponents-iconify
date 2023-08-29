package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dropbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.06 1L0 5.61l4.882 3.908L12 5.123L7.06 1ZM0 13.428l7.06 4.61L12 13.914L4.882 9.52L0 13.43Zm12 .486l4.94 4.124l7.06-4.61l-4.882-3.91L12 13.914ZM24 5.61L16.94 1L12 5.124l7.118 4.395L24 5.609ZM12.014 14.8L7.06 18.913l-2.12-1.385v1.552l7.074 4.243l7.075-4.243v-1.552l-2.12 1.385l-4.955-4.112Z"/>`),
		g.Group(children),
	)
}