package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tablet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M440 16H80a24.028 24.028 0 0 0-24 24v432a24.028 24.028 0 0 0 24 24h360a24.028 24.028 0 0 0 24-24V40a24.028 24.028 0 0 0-24-24Zm-8 448H88v-96h344Zm0-128H88V48h344Z"/><path fill="currentColor" d="M232 400h32v32h-32z"/>`),
		g.Group(children),
	)
}