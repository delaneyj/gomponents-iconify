package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CubeFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTCubeFour0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="m15.34 9l-8.66 5v20l8.66 5L24 44l8.66-5l8.66-5V14l-8.66-5L24 4l-8.66 5Z"/><path d="m24 24l-10.392 6M24 24V11v13Zm0 0l10.392 6L24 24Zm2.815 11.375L24 37l-2.814-1.625m11.258-19.5l2.814 1.625v3.25m-22.516 0V17.5l2.814-1.625"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTCubeFour0)"/>`),
		g.Group(children),
	)
}