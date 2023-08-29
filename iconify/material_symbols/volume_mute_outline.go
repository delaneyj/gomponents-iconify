package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeMuteOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 15V9h4l5-5v16l-5-5H7Zm2-2h2.85L14 15.15v-6.3L11.85 11H9v2Zm2.5-1Z"/>`),
		g.Group(children),
	)
}