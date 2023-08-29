package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Antivirusalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M934 934q-90 90-217 90t-216-90L90 523Q0 434 0 307T90 90T307 0t216 90l411 411q90 89 90 216t-90 217zm-94-365L624 353L353 624l216 216q56 56 135.5 56T840 840t56-135.5T840 569z"/>`),
		g.Group(children),
	)
}