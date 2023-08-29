package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsBh0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsBh0)"><path fill="#eee" d="M0 0h182.5l88.1 268.5l-88 243.5H0z"/><path fill="#d80027" d="m182.5 0l-82.3 42.7l82.3 42.7l-82.3 42.6l82.3 42.7l-82.3 42.7l82.3 42.6l-82.3 42.7l82.3 42.7l-82.3 42.6l82.3 42.7l-82.3 42.7l82.3 42.6H512V0H182.5z"/></g>`),
		g.Group(children),
	)
}