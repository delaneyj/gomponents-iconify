package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerTowerDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M200 40v176a8 8 0 0 1-8 8H64a8 8 0 0 1-8-8V40a8 8 0 0 1 8-8h128a8 8 0 0 1 8 8Z" opacity=".2"/><path d="M88 72a8 8 0 0 1 8-8h64a8 8 0 0 1 0 16H96a8 8 0 0 1-8-8Zm8 40h64a8 8 0 0 0 0-16H96a8 8 0 0 0 0 16Zm112-72v176a16 16 0 0 1-16 16H64a16 16 0 0 1-16-16V40a16 16 0 0 1 16-16h128a16 16 0 0 1 16 16Zm-16 0H64v176h128Zm-64 128a12 12 0 1 0 12 12a12 12 0 0 0-12-12Z"/></g>`),
		g.Group(children),
	)
}