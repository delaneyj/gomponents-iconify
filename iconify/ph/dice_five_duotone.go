package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiceFiveDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M216 64v128a24 24 0 0 1-24 24H64a24 24 0 0 1-24-24V64a24 24 0 0 1 24-24h128a24 24 0 0 1 24 24Z" opacity=".2"/><path d="M192 32H64a32 32 0 0 0-32 32v128a32 32 0 0 0 32 32h128a32 32 0 0 0 32-32V64a32 32 0 0 0-32-32Zm16 160a16 16 0 0 1-16 16H64a16 16 0 0 1-16-16V64a16 16 0 0 1 16-16h128a16 16 0 0 1 16 16ZM104 92a12 12 0 1 1-12-12a12 12 0 0 1 12 12Zm72 0a12 12 0 1 1-12-12a12 12 0 0 1 12 12Zm-72 72a12 12 0 1 1-12-12a12 12 0 0 1 12 12Zm36-36a12 12 0 1 1-12-12a12 12 0 0 1 12 12Zm36 36a12 12 0 1 1-12-12a12 12 0 0 1 12 12Z"/></g>`),
		g.Group(children),
	)
}