package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsGy0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsGy0)"><path fill="#6da544" d="M77.7 0H512v512H77.8z"/><path fill="#eee" d="M425.4 254.7L31.4 512h46.4L512 256L77.7 0H31.4z"/><path fill="#ffda44" d="M256 256L31.4 512l436.8-256L31.4 0z"/><path fill="#333" d="M0 0v1.8l219.6 253.8L0 510v2h31.4l256-256L31.4 0z"/><path fill="#d80027" d="M0 0v512l256-256L0 0z"/></g>`),
		g.Group(children),
	)
}