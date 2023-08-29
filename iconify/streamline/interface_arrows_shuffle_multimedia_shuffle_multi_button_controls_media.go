package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsShuffleMultimediaShuffleMultiButtonControlsMedia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m11.5 2l2 2l-2 2m-3-2h5m-2 4l2 2l-2 2"/><path d="M.5 4H4l4 6h5.5m-13 0H4"/></g>`),
		g.Group(children),
	)
}