package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cocktail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feCocktail0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feCocktail1" fill="currentColor" fill-rule="nonzero"><path id="feCocktail2" d="M13 19h3a1 1 0 0 1 0 2H8a1 1 0 0 1 0-2h3v-6.75L4.8 7.6A2 2 0 0 1 4 6V5a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v1a2 2 0 0 1-.8 1.6L13 12.25V19ZM6 5v1l6 4l6-4V5H6Z"/></g></g>`),
		g.Group(children),
	)
}