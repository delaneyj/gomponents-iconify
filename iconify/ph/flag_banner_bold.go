package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagBannerBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M234.15 49.59A12 12 0 0 0 224 44H32a12 12 0 0 0-8.11 20.85L63 100.71l-39.82 43.15A12 12 0 0 0 32 164h127.28l-26.11 54.84a12 12 0 1 0 21.66 10.32l80-168a12 12 0 0 0-.68-11.57ZM170.71 140H59.41l29.41-31.86a12 12 0 0 0-.71-17L62.85 68H205Z"/>`),
		g.Group(children),
	)
}