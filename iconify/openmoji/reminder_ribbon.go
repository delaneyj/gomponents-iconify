package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReminderRibbon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#D22F27" d="M21.8 62.4L33 66l20-42.7c2.9-6 0-13.4-5.7-15.3"/><path fill="#EA5A47" d="m45 48.5l-6.5 13.1l1.7 3.4l11.2-3.5M25 8c-5.8 1.9-8.6 9.2-5.8 15.3L28 42.2l6.4-13.8"/><path fill="#D22F27" d="m41 20l5.9-12.8H25.3l6 12.8z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="1.813" d="M21.8 62.4L33 66l20-42.7c2.9-6 0-13.4-5.7-15.3L21.8 62.4zm17.1 0l1.3 2.6l11.2-3.5l-6-12.2M34.2 27.6L25 8c-5.8 1.9-8.6 9.2-5.8 15.3l8.5 18.3m6-22.5h4.6m5.4-11H28.6"/>`),
		g.Group(children),
	)
}