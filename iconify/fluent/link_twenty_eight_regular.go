package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11 8a.75.75 0 0 1 .102 1.493L11 9.5H8a4.5 4.5 0 0 0-.212 8.995L8 18.5h3a.75.75 0 0 1 .102 1.493L11 20H8a6 6 0 0 1-.225-11.996L8 8h3Zm9 0a6 6 0 0 1 .225 11.996L20 20h-3a.75.75 0 0 1-.102-1.493L17 18.5h3a4.5 4.5 0 0 0 .212-8.995L20 9.5h-3a.75.75 0 0 1-.102-1.493L17 8h3ZM8 13.25h12a.75.75 0 0 1 .102 1.493L20 14.75H8a.75.75 0 0 1-.102-1.493L8 13.25h12H8Z"/>`),
		g.Group(children),
	)
}