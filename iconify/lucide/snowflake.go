package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Snowflake(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 12h20M12 2v20m8-6l-4-4l4-4M4 8l4 4l-4 4M16 4l-4 4l-4-4m0 16l4-4l4 4"/>`),
		g.Group(children),
	)
}