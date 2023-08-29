package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NameSpace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M7 12H4a2.002 2.002 0 0 1-2-2V6a2.002 2.002 0 0 1 2-2h3v2H4v4h3zm2-2h6v2H9zm8 0h6v2h-6zm11 2h-3v-2h3V6h-3V4h3a2.002 2.002 0 0 1 2 2v4a2.002 2.002 0 0 1-2 2zM17 4h6v2h-6zM9 4h6v2H9zm19 24H4a2.002 2.002 0 0 1-2-2v-4a2.002 2.002 0 0 1 2-2h24a2.002 2.002 0 0 1 2 2v4a2.002 2.002 0 0 1-2 2zM4 22v4h24v-4zm-2-7h28v2H2z"/>`),
		g.Group(children),
	)
}