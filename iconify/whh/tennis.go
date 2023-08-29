package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tennis(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-139 0-257-68.5T68.5 769T0 512q139 0 257-68.5T443.5 257T512 0q139 0 257 68.5T955.5 255t68.5 257q-139 0-257 68.5T580.5 767T512 1024zM4 448q22-174 146-298T448 4q-1 120-60.5 222T226 387.5T4 448zm1016 128q-22 174-146 298t-298 146q1-120 60.5-222T798 636.5t222-60.5z"/>`),
		g.Group(children),
	)
}