package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwoArrowDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8.002 11.854L3.008 8.109v3.94l4.994 3.943l4.965-4.042V8.032l-.012-.011l-4.953 3.833zm0-7.784L3.008.119v3.939l4.994 3.965l4.965-4.064V.041l-.012-.01L8.002 4.07z"/>`),
		g.Group(children),
	)
}