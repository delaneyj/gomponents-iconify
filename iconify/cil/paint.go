package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M418.968 215.116A71.976 71.976 0 0 0 408.023 72H384V40a24.028 24.028 0 0 0-24-24H72a24.028 24.028 0 0 0-24 24v104a24.028 24.028 0 0 0 24 24h288a24.028 24.028 0 0 0 24-24v-40h24.023a39.977 39.977 0 0 1 6.079 79.489L240 210.273V280h-56v216h144V280h-56v-42.273ZM352 136H80V48h272Zm-56 176v152h-80V312Z"/>`),
		g.Group(children),
	)
}