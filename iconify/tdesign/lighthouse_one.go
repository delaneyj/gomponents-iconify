package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LighthouseOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 .923l6.3 2.52l-.743 1.857L17 5.077V10h2v2h-1.907c.164 1.905.52 3.938.885 5.669a71.324 71.324 0 0 0 .962 3.982l.016.057l.004.014v.003L19.326 23H13.22l-1-4h-.439l-1 4H4.674l.364-1.275l.001-.004l.004-.013l.016-.057l.062-.224a71.247 71.247 0 0 0 .9-3.758c.364-1.731.721-3.764.885-5.67H5v-2h2V5.078l-.558.223L5.7 3.443L12 .923ZM9 4.277V10h6V4.277l-3-1.2l-3 1.2ZM15.086 12H8.913c-.168 2.086-.556 4.28-.935 6.08A73.346 73.346 0 0 1 7.3 21h1.92l1-4h3.561l1 4H16.7a73.47 73.47 0 0 1-.678-2.92c-.38-1.8-.767-3.994-.935-6.08ZM11 5.998h2.004v2.004H11V5.998Z"/>`),
		g.Group(children),
	)
}