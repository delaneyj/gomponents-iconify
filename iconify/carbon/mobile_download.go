package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MobileDownload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m28 25l-1.414-1.414L24 26.172V18h-2v8.172l-2.586-2.586L18 25l5 5l5-5z"/><path fill="currentColor" d="M10 28V10h12v5h2V6a2.002 2.002 0 0 0-2-2H10a2.002 2.002 0 0 0-2 2v22a2.002 2.002 0 0 0 2 2h6v-2Zm0-22h12v2H10Z"/>`),
		g.Group(children),
	)
}