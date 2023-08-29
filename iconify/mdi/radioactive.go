package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Radioactive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 11c0 1.1-.9 2-2 2s-2-.9-2-2s.9-2 2-2s2 .9 2 2m2 0h6c0-3.7-2-6.9-5-8.7l-3 5.2c1.2.7 2 2 2 3.5m-4 4c-.7 0-1.4-.2-2-.6c-1.2 2.1-2.6 4.4-3 5.2c1.5.9 3.2 1.3 5 1.3s3.5-.5 5-1.3l-3-5.2c-.6.4-1.3.6-2 .6m-2-7.4C8.8 5.5 7.4 3.1 7 2.3C4 4 2 7.3 2 11h6c0-1.5.8-2.8 2-3.4Z"/>`),
		g.Group(children),
	)
}