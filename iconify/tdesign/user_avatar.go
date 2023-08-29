package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserAvatar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2V2Zm17.5 18h.5V4H4v16h.5a5 5 0 0 1 5-5h5a5 5 0 0 1 5 5ZM12 7a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5ZM7.5 9.5a4.5 4.5 0 1 1 9 0a4.5 4.5 0 0 1-9 0Zm2 7.5a3 3 0 0 0-3 3h11a3 3 0 0 0-3-3h-5Z"/>`),
		g.Group(children),
	)
}