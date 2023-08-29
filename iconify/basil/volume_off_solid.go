package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeOffSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.61 4.846a1 1 0 0 0-1.633-.645L7.958 8.384a.5.5 0 0 1-.32.116H3.75c-.69 0-1.25.56-1.25 1.25v4.5c0 .69.56 1.25 1.25 1.25h3.888a.5.5 0 0 1 .32.116l5.02 4.182a1 1 0 0 0 1.632-.644l.095-.766c.53-4.242.53-8.534 0-12.776l-.095-.766Zm2.627 4.856a.75.75 0 0 1 1.061 0l1.238 1.237l1.237-1.237a.75.75 0 1 1 1.06 1.06L20.597 12l1.238 1.237a.75.75 0 1 1-1.061 1.061l-1.238-1.237l-1.237 1.237a.75.75 0 0 1-1.06-1.06L18.474 12l-1.238-1.238a.75.75 0 0 1 0-1.06Z"/>`),
		g.Group(children),
	)
}