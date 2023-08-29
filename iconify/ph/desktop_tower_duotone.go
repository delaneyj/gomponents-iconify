package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DesktopTowerDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M240 48v160a8 8 0 0 1-8 8h-80a8 8 0 0 1-8-8V48a8 8 0 0 1 8-8h80a8 8 0 0 1 8 8Z" opacity=".2"/><path d="M24 96v72a8 8 0 0 0 8 8h80a8 8 0 0 1 0 16H96v16h16a8 8 0 0 1 0 16H64a8 8 0 0 1 0-16h16v-16H32a24 24 0 0 1-24-24V96a24 24 0 0 1 24-24h80a8 8 0 0 1 0 16H32a8 8 0 0 0-8 8Zm184-32h-32a8 8 0 0 0 0 16h32a8 8 0 0 0 0-16Zm0 32h-32a8 8 0 0 0 0 16h32a8 8 0 0 0 0-16Zm40-48v160a16 16 0 0 1-16 16h-80a16 16 0 0 1-16-16V48a16 16 0 0 1 16-16h80a16 16 0 0 1 16 16Zm-16 160V48h-80v160h80Zm-40-40a12 12 0 1 0 12 12a12 12 0 0 0-12-12Z"/></g>`),
		g.Group(children),
	)
}