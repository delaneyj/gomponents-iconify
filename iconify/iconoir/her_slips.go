package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HerSlips(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="1.5" d="M1 4.6a.6.6 0 0 1 .6-.6h20.8a.6.6 0 0 1 .6.6v3.912c0 .284-.199.53-.476.595c-1.052.247-3.635.914-5.524 1.893c-3.444 1.786-3.93 6.655-3.993 8.382a.637.637 0 0 1-.627.618h-.76a.637.637 0 0 1-.627-.618C10.93 17.655 10.443 12.786 7 11c-1.889-.98-4.472-1.646-5.524-1.893A.614.614 0 0 1 1 8.512V4.6Z"/>`),
		g.Group(children),
	)
}