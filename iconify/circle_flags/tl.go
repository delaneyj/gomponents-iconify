package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsTl0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsTl0)"><path fill="#ffda44" d="m0 0l214 251.8L0 512l418-256z"/><path fill="#d80027" d="M512 0H0l367.3 256L0 512h512z"/><path fill="#333" d="M0 0v512l256-256z"/><path fill="#eee" d="m71 197.4l39 36.8l47-25.6l-23 48.4l39 36.9l-53.2-7l-23 48.5l-9.9-52.7l-53.2-7l47.1-25.6z"/></g>`),
		g.Group(children),
	)
}