package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NinePointsConnected(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m26.829 37.172l10.343-10.343m-16 0L10.829 37.172m16-16l10.343-10.343M10.829 21.172l10.343-10.343M12 8a4 4 0 1 1-8 0a4 4 0 0 1 8 0Zm16 0a4 4 0 1 1-8 0a4 4 0 0 1 8 0Zm16 0a4 4 0 1 1-8 0a4 4 0 0 1 8 0Zm0 16a4 4 0 1 1-8 0a4 4 0 0 1 8 0Zm0 16a4 4 0 1 1-8 0a4 4 0 0 1 8 0Zm-16 0a4 4 0 1 1-8 0a4 4 0 0 1 8 0Zm-16 0a4 4 0 1 1-8 0a4 4 0 0 1 8 0Zm0-16a4 4 0 1 1-8 0a4 4 0 0 1 8 0Zm16 0a4 4 0 1 1-8 0a4 4 0 0 1 8 0Z"/>`),
		g.Group(children),
	)
}