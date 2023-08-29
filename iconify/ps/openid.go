package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Openid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M417 171q-62-36-134-46V6l-70 45v73q-45 4-86 19t-68 37.5t-44 50T2 287t24 55t68.5 48.5T213 426l70-45V173q39 6 90 28l-40 25h129v-85zm-204-2v212q-60-8-96.5-31T75 298.5t7.5-56.5t48-48.5T213 169z"/>`),
		g.Group(children),
	)
}