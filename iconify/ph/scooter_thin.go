package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScooterThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M212 140a31.29 31.29 0 0 0-6.24.62l-11.82-35.46l-22.15-66.42A4 4 0 0 0 168 36h-32a4 4 0 0 0 0 8h29.12l20.54 61.63L134 172H76a32 32 0 1 0-1 8h61a4 4 0 0 0 3.16-1.54l49.54-63.7l9.47 28.39A32 32 0 1 0 212 140ZM44 196a24 24 0 1 1 24-24a24 24 0 0 1-24 24Zm168 0a24 24 0 1 1 24-24a24 24 0 0 1-24 24Z"/>`),
		g.Group(children),
	)
}