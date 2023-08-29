package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReferencesRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15 3v16h5V3zm1 11h3v1h-3zm0-3h3v1h-3zM9 3v16h5V3zm1 11h3v1h-3zm0-3h3v1h-3zM4.1 2.3l-4 15.3l4.8 1.3L9 3.5zM2.3 13.1l2.9.8l-.3 1l-2.9-.8zm.7-2.9l2.9.8l-.2 1l-2.9-.8z"/>`),
		g.Group(children),
	)
}