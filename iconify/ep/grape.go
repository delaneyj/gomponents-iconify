package ep

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Grape(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M544 195.2a160 160 0 0 1 96 60.8a160 160 0 1 1 146.24 254.976a160 160 0 0 1-128 224a160 160 0 1 1-292.48 0a160 160 0 0 1-128-224A160 160 0 1 1 384 256a160 160 0 0 1 96-60.8V128h-64a32 32 0 0 1 0-64h192a32 32 0 0 1 0 64h-64v67.2zM512 448a96 96 0 1 0 0-192a96 96 0 0 0 0 192zm-256 0a96 96 0 1 0 0-192a96 96 0 0 0 0 192zm128 224a96 96 0 1 0 0-192a96 96 0 0 0 0 192zm128 224a96 96 0 1 0 0-192a96 96 0 0 0 0 192zm128-224a96 96 0 1 0 0-192a96 96 0 0 0 0 192zm128-224a96 96 0 1 0 0-192a96 96 0 0 0 0 192z"/>`),
		g.Group(children),
	)
}