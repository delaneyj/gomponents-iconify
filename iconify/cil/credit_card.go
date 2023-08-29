package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M472 72H40a24.028 24.028 0 0 0-24 24v320a24.028 24.028 0 0 0 24 24h432a24.028 24.028 0 0 0 24-24V96a24.028 24.028 0 0 0-24-24Zm-8 32v64H48v-64ZM48 408V232h416v176Z"/><path fill="currentColor" d="M88 312h64v32H88zm96 0h64v32h-64z"/>`),
		g.Group(children),
	)
}