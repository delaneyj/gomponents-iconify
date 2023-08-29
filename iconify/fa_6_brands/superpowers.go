package fa_6_brands

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Superpowers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 448 512"),
		g.Raw(`<path fill="currentColor" d="M448 32c-83.3 11-166.8 22-250 33c-92 12.5-163.3 86.7-169 180c-3.3 55.5 18 109.5 57.8 148.2L0 480c83.3-11 166.5-22 249.8-33c91.8-12.5 163.3-86.8 168.7-179.8c3.5-55.5-18-109.5-57.7-148.2L448 32zm-79.7 232.3c-4.2 79.5-74 139.2-152.8 134.5c-79.5-4.7-140.7-71-136.3-151c4.5-79.2 74.3-139.3 153-134.5c79.3 4.7 140.5 71 136.1 151z"/>`),
		g.Group(children),
	)
}