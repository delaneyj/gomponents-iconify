package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsBg0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsBg0)"><path fill="#496e2d" d="m0 166.9l258-31.7l254 31.7v178l-251.4 41.3L0 344.9z"/><path fill="#eee" d="M0 0h512v166.9H0z"/><path fill="#d80027" d="M0 344.9h512V512H0z"/></g>`),
		g.Group(children),
	)
}