package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortNumbericDesc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M5 12V0H3v12H.5L4 15.5L7.5 12H5zm8.5 4a.5.5 0 0 1-.5-.5V10h-.5a.5.5 0 0 1 0-1h1a.5.5 0 0 1 .5.5v6a.5.5 0 0 1-.5.5zm1-16h-3a.5.5 0 0 0-.5.5v3a.5.5 0 0 0 .5.5H14v2h-2.5a.5.5 0 0 0 0 1h3a.5.5 0 0 0 .5-.5v-6a.5.5 0 0 0-.5-.5zM12 1h2v2h-2V1z"/>`),
		g.Group(children),
	)
}