package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zenphoto(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M832.59 705h-64v160q0 13-9.5 22.5t-22.5 9.5h-64q-13 0-22.5-9.5t-9.5-22.5V353q0-13 9.5-22.5t22.5-9.5h160q80 0 136 56t56 135.5t-56 136t-136 56.5zm0-256h-64v128h64q27 0 45.5-19t18.5-45.5t-18.5-45t-45.5-18.5zm-256-416v-2v2zm-492 800h460q13 0 22.5 9.5t9.5 22.5t-9.5 22.5t-22.5 9.5h-511q-12 2-22.5-8.5T.59 864q-1-15 6-28l486-771h-460q-13 0-22.5-9.5T.59 33t9.5-22.5T32.59 1h511q12-2 22.5 8.5t10.5 23.5q-1 21-6 29zm492-802z"/>`),
		g.Group(children),
	)
}