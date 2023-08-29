package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignRightBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M228 40v176a12 12 0 0 1-24 0V40a12 12 0 0 1 24 0Zm-40 24v36a20 20 0 0 1-20 20H80a20 20 0 0 1-20-20V64a20 20 0 0 1 20-20h88a20 20 0 0 1 20 20Zm-24 4H84v28h80Zm24 88v36a20 20 0 0 1-20 20H40a20 20 0 0 1-20-20v-36a20 20 0 0 1 20-20h128a20 20 0 0 1 20 20Zm-24 4H44v28h120Z"/>`),
		g.Group(children),
	)
}