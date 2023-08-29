package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vectoralt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M576 832q-47 0-88-6l70 70h50q13 0 22.5 9.5T640 928v64q0 13-9.5 22.5T608 1024h-64q-13 0-22.5-9.5T512 992v-50L399 830q-8 2-15 2H256q-27 0-45.5-18.5T192 768V640q0-7 2-15L81 512H32q-13 0-22.5-9.5T0 480v-64q0-13 9.5-22.5T32 384h64q13 0 22.5 9.5T128 416v50l70 70q-6-41-6-88q0-49 23.5-115.5t64-132T382 79T512 0v64q-63 31-122.5 98T293 307t-37 141q0 69 18 128h110q27 0 45.5 19t18.5 45v110q59 18 128 18q63 0 141-37t145-96.5T960 512h64q-23 68-79 130T823.5 744.5t-132 64T576 832zM319.5 640q-26.5 0-45 19T256 704.5t18.5 45t45 18.5t45.5-18.5t19-45t-19-45.5t-45.5-19z"/>`),
		g.Group(children),
	)
}