package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Edge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M.241 7.103C.71 3.403 3.235.047 7.76 0c2.731.053 4.978 1.291 6.316 3.65c.672 1.231.881 2.525.925 3.953v1.678H4.96c.047 4.141 6.094 4 8.697 2.175v3.372c-1.525.916-4.984 1.734-7.662.681c-2.281-.856-3.906-3.244-3.897-5.541c-.075-2.978 1.481-4.95 3.897-6.072c-.513.634-.903 1.334-1.106 2.547h5.669s.331-3.388-3.209-3.388C4.011 3.171 1.605 5.111.243 7.102z"/>`),
		g.Group(children),
	)
}