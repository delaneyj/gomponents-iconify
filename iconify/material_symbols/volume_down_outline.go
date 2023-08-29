package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeDownOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 15V9h4l5-5v16l-5-5H5Zm11 1V7.95q1.125.525 1.813 1.625T18.5 12q0 1.325-.688 2.4T16 16Zm-4-7.15L9.85 11H7v2h2.85L12 15.15v-6.3ZM9.5 12Z"/>`),
		g.Group(children),
	)
}