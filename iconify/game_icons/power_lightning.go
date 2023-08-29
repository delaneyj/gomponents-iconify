package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PowerLightning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M29.805 29.777L242.14 209.55H118.712l112.54 86.784H95.995l225.656 174.012l-81.537-116.05l66.487.143l179.185 138.175l-171.96-244.746h84.568L248.082 29.776H29.805z"/>`),
		g.Group(children),
	)
}