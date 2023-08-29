package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OUpperCase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M788.018 318c37 211-107 412-324 450c-216 38-421-101-458-312c-38-210 108-411 324-449s421 101 458 311zm-391 383c178 0 323-140 323-313s-145-315-323-315s-322 142-322 315s144 313 322 313z"/>`),
		g.Group(children),
	)
}