package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cloud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feCloud0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feCloud1" fill="currentColor" fill-rule="nonzero"><path id="feCloud2" d="m12.713 9.72l-1.073 1.542l-1.606-.975A2 2 0 0 0 7 12v2H5a1 1 0 0 0 0 2h11a4 4 0 1 0-3.287-6.28ZM16 6a6 6 0 1 1 0 12H5a3 3 0 0 1 0-6a4 4 0 0 1 6.071-3.423A5.993 5.993 0 0 1 16 6Z"/></g></g>`),
		g.Group(children),
	)
}