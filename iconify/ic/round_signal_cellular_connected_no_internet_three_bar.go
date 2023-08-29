package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundSignalCellularConnectedNoInternetThreeBar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-opacity=".3" d="M22 8V4.41c0-.89-1.08-1.34-1.71-.71L3.71 20.29c-.63.63-.19 1.71.7 1.71H18V11c0-1.66 1.34-3 3-3h1z"/><path fill="currentColor" d="M18 22V6L3.71 20.29c-.63.63-.19 1.71.7 1.71H18zm2-11v6c0 .55.45 1 1 1s1-.45 1-1v-6c0-.55-.45-1-1-1s-1 .45-1 1zm0 11h2v-2h-2v2z"/>`),
		g.Group(children),
	)
}