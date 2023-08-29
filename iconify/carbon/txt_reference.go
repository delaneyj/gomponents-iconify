package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TxtReference(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4 20v2h4.586L2 28.586L3.414 30L10 23.414V28h2v-8H4zM21 4h3v12h2V4h3V2h-8v2zm-1-2h-2l-2 6l-2-6h-2l2.752 7L12 16h2l2-6l2 6h2l-2.755-7L20 2zM3 4h3v12h2V4h3V2H3v2z"/>`),
		g.Group(children),
	)
}