package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 10h-5V6a2.002 2.002 0 0 0-2-2H11a2.002 2.002 0 0 0-2 2v4H4a2.002 2.002 0 0 0-2 2v16a2.002 2.002 0 0 0 2 2h24a2.002 2.002 0 0 0 2-2V12a2.002 2.002 0 0 0-2-2ZM4 28V12h5v2H7v2h2v2H7v2h2v2H7v2h2v4Zm17 0H11V6h10Zm7 0h-5v-4h2v-2h-2v-2h2v-2h-2v-2h2v-2h-2v-2h5Z"/><path fill="currentColor" d="M14 8h4v2h-4zm0 4h4v2h-4zm0 4h4v2h-4z"/>`),
		g.Group(children),
	)
}