package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Discourse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.102 0h-.081C5.462 0 .13 5.252.001 11.779v.012L.007 24l12.097-.01c6.582-.055 11.897-5.404 11.897-11.995S18.686.056 12.109 0h-.005zM12 18.857h-.015a6.778 6.778 0 0 1-2.94-.666l.041.018l-4.345 1.077l1.227-4.018a6.78 6.78 0 0 1-.83-3.27A6.86 6.86 0 1 1 12 18.857z"/>`),
		g.Group(children),
	)
}