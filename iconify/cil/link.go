package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Link(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M457.47 55.833c-53.026-53.026-139.307-53.026-192.332 0L168.971 152l22.629 22.627l96.165-96.167a104 104 0 0 1 147.078 147.079l-96.167 96.167l22.624 22.627l96.167-96.167C510.5 195.14 510.5 108.86 457.47 55.833Zm-231.931 379.01a104 104 0 0 1-147.078 0a104 104 0 0 1 0-147.078l90.511-90.511l-22.627-22.627l-90.512 90.511A136 136 0 1 0 248.166 457.47l90.51-90.51l-22.627-22.627Z"/><path fill="currentColor" d="m129.373 361.303l226.274-226.275l22.628 22.628L152 383.93z"/>`),
		g.Group(children),
	)
}