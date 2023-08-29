package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ve(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsVe0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsVe0)"><path fill="#0052b4" d="m0 144.7l255.3-36.5L512 144.7v222.6L250.5 407L0 367.3z"/><path fill="#ffda44" d="M0 0h512v144.7H0z"/><path fill="#d80027" d="M0 367.3h512V512H0z"/><path fill="#eee" d="M443.4 306.4L429.8 317l6 16.1l-14.3-9.6l-13.5 10.7l4.7-16.5l-14.2-9.6l17.1-.6l4.7-16.5l6 16.1zm-34.7-60l-9 14.5l11 13.2L394 270l-9 14.6l-1.3-17l-16.6-4.3l15.9-6.4l-1.2-17l11 13zm-53-44.5l-3.6 16.8l14.9 8.4l-17 1.8l-3.6 16.8l-7-15.7l-17 1.8l12.7-11.5l-7-15.6l14.8 8.6zm-65-23.7l2.3 17l17 3l-15.5 7.5l2.4 17l-12-12.4l-15.4 7.6l8-15.2l-11.8-12.3l16.9 3zm-69.3 0l8 15.1l17-3l-12 12.4l8 15.2l-15.4-7.6l-11.9 12.4l2.4-17l-15.4-7.5l16.9-3zm-65 23.7l12.7 11.5l14.8-8.6l-7 15.7l12.8 11.4l-17-1.8l-7 15.7l-3.7-16.7l-17-1.8l14.8-8.5zm-53.1 44.5l15.9 6.4l11-13l-1.2 17l16 6.4l-16.7 4.2l-1.2 17L118 270l-16.7 4.2l11-13.2zm-34.7 60l17.2.5l5.9-16l4.7 16.4l17.1.6l-14.2 9.6l4.7 16.6l-13.5-10.6l-14.2 9.6l5.9-16z"/></g>`),
		g.Group(children),
	)
}