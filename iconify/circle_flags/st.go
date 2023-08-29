package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func St(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsSt0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsSt0)"><path fill="#6da544" d="M0 0h512v167l-52.6 83.8L512 345v167H0l72-264.3z"/><path fill="#ffda44" d="M114.9 166.9H512v178H114.9z"/><path fill="#d80027" d="M0 0v512l256-256z"/><path fill="#333" d="m325 211.5l11.1 34H372l-29 21l11.1 34l-29-21l-28.9 21l11-34l-28.8-21H314zm111.4 0l11 34h35.8l-29 21l11.1 34l-29-21l-28.9 21l11.1-34l-29-21h35.8z"/></g>`),
		g.Group(children),
	)
}