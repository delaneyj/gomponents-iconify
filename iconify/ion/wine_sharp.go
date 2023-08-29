package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M453 112V66.33H60.75V112l175.13 176v118H124.75v42H389v-42H277.88V288Zm-336.65-3.67h281l-37.81 38H154.16Z"/>`),
		g.Group(children),
	)
}