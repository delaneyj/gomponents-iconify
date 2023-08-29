package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M416 174.74V48h-80v58.45L256 32L0 272h64v208h144V320h96v160h144V272h64l-96-97.26z"/>`),
		g.Group(children),
	)
}