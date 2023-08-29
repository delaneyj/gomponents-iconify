package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DsVideo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 5.5h37v37h-37z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16 8.5h16v14H16zm0 17h16v14H16zm19.5-8h4v4h-4zm0-8h4v4h-4zm0 17h4v4h-4zm0 8h4v4h-4zm-26.887-17h4v4h-4zm0-8h4v4h-4zm0 17h4v4h-4zm0 8h4v4h-4z"/>`),
		g.Group(children),
	)
}