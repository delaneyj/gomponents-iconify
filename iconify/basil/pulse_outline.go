package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PulseOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8.24 4.876a.75.75 0 0 0-1.454-.106l-2.09 6.48H2a.75.75 0 0 0 0 1.5h3.242a.75.75 0 0 0 .714-.52L7.27 8.16l1.84 10.965a.75.75 0 0 0 1.46.083l2.43-8.48l1.28 4.48a.75.75 0 0 0 1.368.172l1.54-2.629h1.166a2.751 2.751 0 1 0 0-1.5h-1.595a.75.75 0 0 0-.648.37l-.87 1.488l-1.519-5.314a.75.75 0 0 0-1.442 0l-2.239 7.814l-1.8-10.732ZM19.75 12a1.25 1.25 0 1 1 2.5 0a1.25 1.25 0 0 1-2.5 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}