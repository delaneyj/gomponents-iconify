package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ResizeBoth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m404.687 212.686l-50.51 50.51l-105.372-105.373l50.841-50.843L272.666 80H80v192l27.313 27.314l50.982-50.981l105.372 105.373l-51.313 51.313L239.333 432H432V240ZM400 400H262.627l46.3-46.294l-150.632-150.628L112 249.373V112h137.373l-45.823 45.823l150.627 150.628L400 262.627Z"/>`),
		g.Group(children),
	)
}