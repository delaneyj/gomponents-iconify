package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ParkingOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21.933A9.933 9.933 0 1 1 21.933 12A9.944 9.944 0 0 1 12 21.933Zm0-18.866A8.933 8.933 0 1 0 20.933 12A8.943 8.943 0 0 0 12 3.067Z"/><path fill="currentColor" d="M12.569 8.5h-1.75a.749.749 0 0 0-.75.75v5.74a.5.5 0 0 0 .5.5a.5.5 0 0 0 .5-.5V13.5h1.5a2.5 2.5 0 0 0 0-5Zm0 4h-1.5v-3h1.5a1.5 1.5 0 0 1 0 3Z"/>`),
		g.Group(children),
	)
}