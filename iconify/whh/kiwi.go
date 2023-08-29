package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kiwi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M832 691v109q0 13-9.5 22.5T800 832h-64q-13 0-22.5-9.5T704 800t9.5-22.5T736 768h32v-64q-21 0-36.5-17.5T700 635q-14 47-60 63v102q0 13-9.5 22.5T608 832h-64q-13 0-22.5-9.5T512 800t9.5-22.5T544 768h32v-87q-4-8-10.5-27.5t-15-36.5t-20.5-28q-148-50-255-189q-74 48-147 48l-6-1q-31 52-44.5 93T64 640q0 13-11 38.5T32 704q-9 0-20.5-24.5T0 640q0-174 66-312q-2-26-2-72q0-53 37.5-90.5T192 128q32 0 73 14q59-66 153-104T608 0q59 0 118 15t112.5 45t95 71t66 98.5T1024 352q0 118-56 212.5T832 691z"/>`),
		g.Group(children),
	)
}