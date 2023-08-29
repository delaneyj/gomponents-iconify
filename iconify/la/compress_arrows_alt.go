package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CompressArrowsAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4.719 3.281L3.28 4.72L10.562 12H5v2h9V5h-2v5.563zm22.562 0L20 10.562V5h-2v9h9v-2h-5.563l7.282-7.281zM5 18v2h5.563L3.28 27.281l1.44 1.439L12 21.437V27h2v-9zm13 0v9h2v-5.563l7.281 7.282l1.438-1.438L21.437 20H27v-2z"/>`),
		g.Group(children),
	)
}