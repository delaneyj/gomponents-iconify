package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LyOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagLy1x10"><path d="M250 12h500v500H250z"/></clipPath></defs><g clip-path="url(#flagLy1x10)" transform="translate(-256 -12.3) scale(1.024)"><path fill="#239e46" d="M0 12h1000v500H0z"/><path d="M0 12h1000v375H0z"/><path fill="#e70013" d="M0 12h1000v125H0z"/><path fill="#fff" d="M544.2 217.8a54.3 54.3 0 1 0 0 88.4a62.5 62.5 0 1 1 0-88.4M530.4 262l84.1-27.3l-52 71.5v-88.4l52 71.5z"/></g>`),
		g.Group(children),
	)
}