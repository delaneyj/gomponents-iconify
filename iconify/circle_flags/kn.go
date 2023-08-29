package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsKn0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsKn0)"><path fill="#ffda44" d="m0 401.9l173.6-225.3L401.9 0H449l63 63l-.1 47.3l-167.3 223.5L110.3 512H63L0 449z"/><path fill="#6da544" d="M0 0v401.9L401.9 0z"/><path fill="#d80027" d="M512 512V110.3L110.3 512z"/><path fill="#333" d="M0 512h63L512 63V0h-63L0 449z"/><path fill="#eee" d="m162.8 302l24 12.2l19-19l-4.4 26.5l24 12.2l-26.6 4.2l-4.2 26.5l-12-24L156 345l19-19zM302 162.8l24 12.1l19-19l-4.3 26.6l24 12.1l-26.6 4.2l-4.2 26.5l-12.2-23.9l-26.5 4.2l19-19z"/></g>`),
		g.Group(children),
	)
}