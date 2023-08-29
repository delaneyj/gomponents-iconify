package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CompassOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.661 20.937a1.589 1.589 0 0 1-1.117-.48a1.534 1.534 0 0 1-.4-1.59l3.436-9.93A3.8 3.8 0 0 1 8.938 6.58l9.93-3.439a1.537 1.537 0 0 1 1.589.4a1.532 1.532 0 0 1 .4 1.588l-3.437 9.932a3.8 3.8 0 0 1-2.358 2.358l-9.93 3.439a1.442 1.442 0 0 1-.471.079ZM19.337 4.062a.424.424 0 0 0-.142.024L9.267 7.525a2.8 2.8 0 0 0-1.742 1.741L4.087 19.2a.6.6 0 0 0 .717.718l9.93-3.439a2.8 2.8 0 0 0 1.741-1.741L19.913 4.8a.551.551 0 0 0-.163-.553a.609.609 0 0 0-.413-.185Z"/><circle cx="12" cy="12" r="1.563" fill="currentColor"/>`),
		g.Group(children),
	)
}