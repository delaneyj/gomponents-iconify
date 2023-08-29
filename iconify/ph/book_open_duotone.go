package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookOpenDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M232 64v128a8 8 0 0 1-8 8h-64a32 32 0 0 0-32 32a32 32 0 0 0-32-32H32a8 8 0 0 1-8-8V64a8 8 0 0 1 8-8h64a32 32 0 0 1 32 32a32 32 0 0 1 32-32h64a8 8 0 0 1 8 8Z" opacity=".2"/><path d="M224 48h-64a40 40 0 0 0-32 16a40 40 0 0 0-32-16H32a16 16 0 0 0-16 16v128a16 16 0 0 0 16 16h64a24 24 0 0 1 24 24a8 8 0 0 0 16 0a24 24 0 0 1 24-24h64a16 16 0 0 0 16-16V64a16 16 0 0 0-16-16ZM96 192H32V64h64a24 24 0 0 1 24 24v112a39.81 39.81 0 0 0-24-8Zm128 0h-64a39.81 39.81 0 0 0-24 8V88a24 24 0 0 1 24-24h64Z"/></g>`),
		g.Group(children),
	)
}