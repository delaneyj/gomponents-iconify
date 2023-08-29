package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VcFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd"><path fill="#f4f100" d="M0 0h640v480H0z"/><path fill="#199a00" d="M490 0h150v480H490z"/><path fill="#0058aa" d="M0 0h150v480H0z"/><path fill="#199a00" d="m259.3 130l-46.4 71.3l44.7 74.4l43.8-73.7l-42.1-72zm121.2 0l-46.3 71.3l44.7 74.4l43.8-73.7l-42.2-72zm-61.2 97.3l-46.4 71.4l44.8 74.4l43.8-73.7l-42.2-72z"/></g>`),
		g.Group(children),
	)
}