package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Repo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14 10V1.5l-.5-.5H3.74a1.9 1.9 0 0 0-.67.13a1.77 1.77 0 0 0-.94 1a1.7 1.7 0 0 0-.13.62v9.5a1.7 1.7 0 0 0 .13.67c.177.427.515.768.94.95a1.9 1.9 0 0 0 .67.13H4v-1h-.26a.72.72 0 0 1-.29-.06a.74.74 0 0 1-.4-.4a.93.93 0 0 1-.05-.29v-.5a.93.93 0 0 1 .05-.29a.74.74 0 0 1 .4-.4a.72.72 0 0 1 .286-.06H13v2H9v1h4.5l.5-.5V10zM4 10V2h9v8H4zm1-7h1v1H5V3zm0 2h1v1H5V5zm1 2H5v1h1V7zm.5 6.49L5.28 15H5v-3h3v3h-.28L6.5 13.49z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}