package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M296 40h32v32h-32zm72 0h32v32h-32zm-72 400h32v32h-32zm72 0h32v32h-32zm72-400h32v32h-32zm0 65.454h32v33.454h-32zm0 66.909h32v33.454h-32zm0 200.728h32v33.454h-32zm0-133.819h32v33.454h-32zm0 66.91h32v33.454h-32zM440 440h32v32h-32zM40 456a16 16 0 0 0 16 16h168v24h32V16h-32v24H56a16 16 0 0 0-16 16ZM72 72h152v368H72Z"/>`),
		g.Group(children),
	)
}