package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Im(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsIm0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsIm0)"><path fill="#d80027" d="M0 0h512v512H0z"/><path fill="#eee" d="m350.8 171.6l-18.1 64.6l-54.3-10l-35-72l-94.4 33.4l-7.4-21l-24.7-3l18.6 52.5l65-16.7l18.4 52l-44.9 66.3l76.3 65l-14.5 17l9.7 22.9l36.1-42.3l-46.8-48l35.8-42l79.8 5.8l18.2-98.6l22 4l15-19.8l-54.8-10zM256 256z"/></g>`),
		g.Group(children),
	)
}