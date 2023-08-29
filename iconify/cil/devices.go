package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Devices(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M472 232h-48V120a24.028 24.028 0 0 0-24-24H40a24.028 24.028 0 0 0-24 24v246a24.028 24.028 0 0 0 24 24h172v50h-60v32h152v-32h-60v-50h92v58a24.027 24.027 0 0 0 24 24h112a24.027 24.027 0 0 0 24-24V256a24.027 24.027 0 0 0-24-24Zm-136 24v102H48V128h344v104h-32a24.027 24.027 0 0 0-24 24Zm128 184h-96V264h96Z"/>`),
		g.Group(children),
	)
}