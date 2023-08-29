package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MobileLandscape(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M464 96H48a32.036 32.036 0 0 0-32 32v256a32.036 32.036 0 0 0 32 32h416a32.036 32.036 0 0 0 32-32V128a32.036 32.036 0 0 0-32-32ZM48 384V128h48v256.018H48Zm80-256h256v256l-256 .013Zm336 256h-48V128h48Z"/>`),
		g.Group(children),
	)
}