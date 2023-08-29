package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paintbrush(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m1015 971l-44 44q-10 9-86.5 9t-86.5-9L503 619l-68 63q-25 23-82 21.5T270 678q-16-15-14-36.5t11-39t20-27.5l288-288q10-9 27-17.5t39.5-12T680 269q25 24 23 95t-27 95l-53 48l392 292q9 9 9 85.5t-9 86.5zM270 678l.5.5l.5.5zm411-409h-1v-1zM393 54q-9-9-9-22t9-22.5T415 0t22 10l184 184l-44 44zm88 280L297 150q-9-10-9-22.5t9-22t22-9.5t22 10l184 184zm-96 96L201 246q-9-10-9-22.5t9-22t22-9.5t22 9l184 185zm-96 95L105 341q-9-9-9-21.5t9-22t22-9.5t22 9l184 184zm-96 96L9 437q-9-9-9-22t9-22t22-9t22 9l184 184z"/>`),
		g.Group(children),
	)
}