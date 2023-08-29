package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Certificatealt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m673 384l96 103l-135 42l32 137l-127-29l102 325q0 33-28 52l-52-52l-62 62q-22-5-36-22.5T449 962l-65-206l-64 206q0 22-14 39.5t-36 22.5l-62-62l-52 52q-28-19-28-52l102-325l-127 29l32-137L0 487l96-103L0 281l135-41l-32-137l137 32L281 0l103 96L487 0l42 135l137-32l-32 137l135 41zM384.5 192q-79.5 0-136 56.5t-56.5 136t56.5 136t136 56.5t136-56.5t56.5-136t-56.5-136t-136-56.5z"/>`),
		g.Group(children),
	)
}