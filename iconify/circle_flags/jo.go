package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Jo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsJo0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsJo0)"><path fill="#eee" d="m126 158l127.8-10.3L512 167v178l-254.9 32.3L126 335.9z"/><path fill="#333" d="M0 0h512v167H107z"/><path fill="#6da544" d="M107 345h405v167H0z"/><path fill="#d80027" d="M0 0v512l256-256z"/><path fill="#eee" d="m101.6 200.3l14 29.4l31.8-7.3l-14.2 29.3l25.5 20.2l-31.8 7.2l.1 32.6l-25.4-20.4l-25.4 20.4V279l-31.7-7.2l25.5-20l-14.2-29.4l31.8 7.3z"/></g>`),
		g.Group(children),
	)
}