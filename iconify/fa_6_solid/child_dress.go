package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChildDress(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M224 64a64 64 0 1 0-128 0a64 64 0 1 0 128 0zM88 400v80c0 17.7 14.3 32 32 32s32-14.3 32-32v-80h16v80c0 17.7 14.3 32 32 32s32-14.3 32-32v-80h17.8c10.9 0 18.6-10.7 15.2-21.1l-31.1-93.4l28.6 37.8c10.7 14.1 30.8 16.8 44.8 6.2s16.8-30.7 6.2-44.8L254.6 207c-22.4-29.6-57.5-47-94.6-47s-72.2 17.4-94.6 47L6.5 284.7c-10.7 14.1-7.9 34.2 6.2 44.8s34.2 7.9 44.8-6.2l28.7-37.8L55 378.9c-3.4 10.4 4.3 21.1 15.2 21.1H88z"/>`),
		g.Group(children),
	)
}