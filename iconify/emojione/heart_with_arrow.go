package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartWithArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#ffce31" d="m12.7 54.6l-2-2.1l19.1-19l2.1 2z"/><path fill="#42ade2" d="M4.2 61s-.7-8.3 6.4-15.3l2.1 6.9L4.2 61"/><path fill="#467591" d="M4.2 61s8.3.7 15.4-6.4l-6.9-2.1L4.2 61"/><path fill="#ff5a79" d="M54.7 24.3c-5.7-15-24.2-8.3-26-.8c-2.4-8-20.4-14-26 .8c-6.1 16.4 23.8 31.2 26 33.6c2.2-1.9 32.2-17.4 26-33.6"/><path fill="#ffce31" d="m37.1 30.1l-2.4-2.3l16.8-16.7l2.3 2.4z"/><path fill="#467591" d="M40.9 21.7s-.4-9.9 8.4-18.7l2.2 8.1l-10.6 10.6m2.3 2.3s9.9.4 18.8-8.4l-8.2-2.2L43.2 24"/>`),
		g.Group(children),
	)
}