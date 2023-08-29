package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Opa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M59.385 242.467c2.672-50.603 44.751-77.785 94.423-54.037c44.02-135.362 263.243-133.97 275.596 49.587c106.336 15.459 113.836 156.376.055 181.12H78.124c-96.737-20.182-104.393-141.973-18.739-176.67z"/>`),
		g.Group(children),
	)
}