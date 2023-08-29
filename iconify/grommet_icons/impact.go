package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Impact(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="m13 2l9 4v11l-9 5V2Zm9 4l-9 5l9-5ZM9 22V2v20Zm0-10L3 5l6 7Zm0 0H1h8Zm0 0l-6 7l6-7Z"/>`),
		g.Group(children),
	)
}