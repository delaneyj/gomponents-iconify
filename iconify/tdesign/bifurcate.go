package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bifurcate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 4a1 1 0 1 0 0 2a1 1 0 0 0 0-2ZM3 5a3 3 0 1 1 4.086 2.797c.128.667.412 1.506.934 2.256C8.752 11.103 9.958 12 12 12s3.248-.897 3.98-1.947a6.01 6.01 0 0 0 .934-2.256a3.001 3.001 0 1 1 2.019.055a8.014 8.014 0 0 1-1.313 3.345c-.933 1.34-2.42 2.478-4.62 2.744v2.23a3.001 3.001 0 1 1-2 0v-2.23c-2.2-.266-3.687-1.405-4.62-2.744a8.014 8.014 0 0 1-1.313-3.345A3.001 3.001 0 0 1 3 5Zm15-1a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm-6 14a1 1 0 1 0 0 2a1 1 0 0 0 0-2Z"/>`),
		g.Group(children),
	)
}