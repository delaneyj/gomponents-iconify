package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Xk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsXk0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsXk0)"><path fill="#0052b4" d="M0 0h512v512H0z"/><path fill="#eee" d="m112.4 155.8l5.6 17h17.9l-14.5 10.5l5.5 17l-14.5-10.5L98 200.3l5.5-17L89 172.8h18zm55.7-16.7l5.5 17h18L177 166.6l5.6 17l-14.5-10.5l-14.5 10.6l5.6-17l-14.5-10.6h17.9zm55.7-16.7l5.5 17h17.9L232.7 150l5.5 17l-14.4-10.6l-14.5 10.6l5.5-17l-14.5-10.6h18zm175.8 33.4l-5.6 17h-17.9l14.5 10.5l-5.5 17l14.5-10.5l14.4 10.5l-5.5-17l14.5-10.5h-18zm-55.7-16.7l-5.5 17h-18l14.6 10.5l-5.6 17l14.5-10.5l14.5 10.6l-5.6-17l14.5-10.6h-17.9zm-55.7-16.7l-5.5 17h-17.9l14.5 10.6l-5.5 17l14.4-10.6l14.5 10.6l-5.5-17l14.5-10.6h-18z"/><path fill="#ffda44" d="M300.5 267.1L256 211.5l-22.3 11.1V245l-33.4 22h-22.2v29a89 89 0 0 1 55.6 82.5H256v-22.2l22.3-22.3l22.2 22.3l22.3-22.3v-22.2l22.2-33.4l-44.5-11.2z"/></g>`),
		g.Group(children),
	)
}