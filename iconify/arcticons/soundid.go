package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Soundid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.71 20.216h26.579v7.568H10.71zm5.432 17.933l10.364-10.364l5.351 5.351L21.493 43.5zm0-23.284L26.506 4.5l5.351 5.351l-10.364 10.364z"/>`),
		g.Group(children),
	)
}