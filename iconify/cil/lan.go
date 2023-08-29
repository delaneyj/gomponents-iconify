package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M496 272v-32H272v-48h56a24.027 24.027 0 0 0 24-24V40a24.027 24.027 0 0 0-24-24H184a24.027 24.027 0 0 0-24 24v128a24.027 24.027 0 0 0 24 24h56v48H16v32h80v48H41.391a24.028 24.028 0 0 0-24 24v128a24.028 24.028 0 0 0 24 24H184a24.027 24.027 0 0 0 24-24V344a24.027 24.027 0 0 0-24-24h-56v-48h256v48h-56a24.027 24.027 0 0 0-24 24v128a24.027 24.027 0 0 0 24 24h144a24.027 24.027 0 0 0 24-24V344a24.027 24.027 0 0 0-24-24h-56v-48ZM192 48h128v112H192Zm-16 416H49.391V352H176Zm288 0H336V352h128Z"/>`),
		g.Group(children),
	)
}