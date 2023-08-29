package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForSolomonIslands(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#699635" d="m60.4 22.5l-8.2-2.2l-40.5 23.4l-2.2 8.2c3.8 4.3 8.8 7.5 14.7 9.1c16 4.3 32.5-5.2 36.7-21.2c1.7-5.9 1.4-11.9-.5-17.3"/><path fill="#2a5f9e" d="M39.8 3C23.8-1.3 7.3 8.2 3 24.2c-1.6 5.9-1.3 11.9.5 17.3l3.7 1l4.5 1.2l40.5-23.4l1.2-4.5l1-3.7C50.7 7.8 45.7 4.6 39.8 3"/><path fill="#ffe62e" d="M3.5 41.5C4.2 43.4 5 45.2 6 47c1 1.7 2.2 3.4 3.5 4.9l50.9-29.4c-1.3-3.8-3.3-7.4-6-10.4L3.5 41.5"/><path fill="#fff" d="m19 20.5l2 1.4l-.8-2.2l2-1.4h-2.4l-.8-2.2l-.8 2.2h-2.4l2 1.4l-.8 2.2zm-5.4-5.7l2 1.3l-.8-2.2l2-1.3h-2.4l-.8-2.2l-.8 2.2h-2.4l2 1.3l-.8 2.2zm10.8 0l2 1.3l-.8-2.2l2-1.3h-2.4l-.8-2.2l-.8 2.2h-2.5l2 1.3l-.7 2.2zM13.6 26.3l2 1.3l-.8-2.2l2-1.3h-2.4l-.8-2.2l-.8 2.2h-2.4l2 1.3l-.8 2.2zm10.8 0l2 1.3l-.8-2.2l2-1.3h-2.4l-.8-2.2l-.8 2.2h-2.5l2 1.3l-.7 2.2z"/>`),
		g.Group(children),
	)
}