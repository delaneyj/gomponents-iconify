package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Turret(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#195DE6" d="m208 288l-16-128h32l16-16V0h-48v48h-32V0H96v48H64V0H16v144l16 16h32L48 288H16L0 304v80h256v-80l-16-16z"/>`),
		g.Group(children),
	)
}