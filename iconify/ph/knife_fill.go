package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KnifeFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M231.81 32.19a28 28 0 0 0-39.67.07L18.27 210.6A8 8 0 0 0 22.2 224a154.93 154.93 0 0 0 35 4c33.42 0 66.88-10.88 98.33-32.21c31.75-21.53 50.15-45.85 50.92-46.88a8 8 0 0 0-.74-10.46l-18.74-18.76l45-48a28.08 28.08 0 0 0-.16-39.5Zm-42.59 112.44a225.51 225.51 0 0 1-43.11 38.18c-34.47 23.25-70 32.7-105.84 28.16l106.3-109Z"/>`),
		g.Group(children),
	)
}