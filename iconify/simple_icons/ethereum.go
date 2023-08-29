package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ethereum(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.944 17.97L4.58 13.62L11.943 24l7.37-10.38l-7.372 4.35h.003zM12.056 0L4.69 12.223l7.365 4.354l7.365-4.35L12.056 0z"/>`),
		g.Group(children),
	)
}