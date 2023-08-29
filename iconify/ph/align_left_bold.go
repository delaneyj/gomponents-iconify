package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignLeftBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M52 40v176a12 12 0 0 1-24 0V40a12 12 0 0 1 24 0Zm16 60V64a20 20 0 0 1 20-20h88a20 20 0 0 1 20 20v36a20 20 0 0 1-20 20H88a20 20 0 0 1-20-20Zm24-4h80V68H92Zm144 60v36a20 20 0 0 1-20 20H88a20 20 0 0 1-20-20v-36a20 20 0 0 1 20-20h128a20 20 0 0 1 20 20Zm-24 4H92v28h120Z"/>`),
		g.Group(children),
	)
}