package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArticleRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 17a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V3a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2zM15 5h-5V4h5zm0 2h-5V6h5zm0 2h-5V8h5zM5 14h10v1H5zm0-2h10v1H5zm0-2h10v1H5zm0-6h4v5H5z"/>`),
		g.Group(children),
	)
}