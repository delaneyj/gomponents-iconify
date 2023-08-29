package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Compress(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feCompress0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feCompress1" fill="currentColor"><path id="feCompress2" d="M18 12h-6V6l2 2l3-3l2 2l-3 3l2 2ZM6 12h6v6l-2-2l-3 3l-2-2l3-3l-2-2Z"/></g></g>`),
		g.Group(children),
	)
}