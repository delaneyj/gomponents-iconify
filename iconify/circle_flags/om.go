package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Om(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsOm0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsOm0)"><path fill="#eee" d="M189.2 0H512v167l-347.5 24.6z"/><path fill="#6da544" d="m163 320l349 25v167H189.2z"/><path fill="#d80027" d="M0 0h189.2v167H512v178H189.2v167H0z"/><path fill="#eee" d="M156.6 112.7L133 89l-15.7 15.8L101.5 89L78 112.7l15.8 15.7L78 144l23.6 23.6l15.8-15.7l15.7 15.7l23.6-23.6l-15.7-15.7z"/></g>`),
		g.Group(children),
	)
}