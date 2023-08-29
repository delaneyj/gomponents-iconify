package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func View(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M397.25 278c38 0 69 31 69 69s-31 68-69 68s-68-30-68-68s30-69 68-69zm0-170c226 0 389 212 389 212c11 14 11 39 0 53c0 0-163 212-389 212s-389-212-389-212c-11-14-11-39 0-53c0 0 163-212 389-212zm0 410c94 0 171-77 171-171s-77-171-171-171s-171 77-171 171s77 171 171 171z"/>`),
		g.Group(children),
	)
}