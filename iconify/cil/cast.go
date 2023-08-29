package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cast(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M16 360v64h64a64 64 0 0 0-64-64ZM472 80H40a24.028 24.028 0 0 0-24 24v72h32v-64h416v280H264v32h208a24.028 24.028 0 0 0 24-24V104a24.028 24.028 0 0 0-24-24Z"/><path fill="currentColor" d="M16 216v32c97.047 0 176 78.953 176 176h32c0-114.691-93.309-208-208-208Z"/><path fill="currentColor" d="M16 288v32a104.118 104.118 0 0 1 104 104h32c0-74.991-61.009-136-136-136Z"/>`),
		g.Group(children),
	)
}