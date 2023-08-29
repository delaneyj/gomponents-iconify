package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Recycling(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M11 20h11.5v-.5l-3.736-6.743M11 20c1.933 0 3.5-3.5 3.5-3.5M11 20c1.933 0 3.5 3.5 3.5 3.5M17.79 11L12.25 1h-.5L8.025 7.73M17.789 11c-.966-1.674-4.781-1.281-4.781-1.281M17.789 11c-.88-1.657 1.281-4.781 1.281-4.781M7.05 9.484L1.5 19.5v.5H9M7.05 9.484c-.88 1.657 1.282 4.781 1.282 4.781M7.05 9.485c-.967 1.673-4.781 1.28-4.781 1.28"/>`),
		g.Group(children),
	)
}