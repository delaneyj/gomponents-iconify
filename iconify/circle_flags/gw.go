package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gw(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsGw0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsGw0)"><path fill="#d80027" d="M0 0h189.2l54 257.6l-54 254.4H0z"/><path fill="#ffda44" d="M189.2 0H512v256l-159 53.5L189.1 256z"/><path fill="#6da544" d="M189.2 256H512v256H189.2z"/><path fill="#333" d="m96.7 189.2l16.6 51H167l-43.4 31.6l16.5 51l-43.4-31.5l-43.4 31.5l16.6-51l-43.4-31.6h53.7z"/></g>`),
		g.Group(children),
	)
}