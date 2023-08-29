package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MdBrowsers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M424 64H88c-26.6 0-48 21.6-48 48v288c0 26.4 21.4 48 48 48h336c26.4 0 48-21.6 48-48V112c0-26.4-21.4-48-48-48zm0 336H88V176h336v224z" fill="currentColor"/>`),
		g.Group(children),
	)
}