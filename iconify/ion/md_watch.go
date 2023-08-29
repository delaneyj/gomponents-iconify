package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MdWatch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M416 256c0-51.001-24.004-96.001-60.996-125L336 16H176l-19.004 115C120.004 159.999 96 204.999 96 256c0 50.996 24.004 95.996 60.996 125L176 496h160l19.004-115C391.996 351.996 416 306.996 416 256zm-280 0c0-66.001 54.004-120 120-120s120 53.999 120 120c0 65.996-54.004 120-120 120s-120-54.004-120-120z" fill="currentColor"/>`),
		g.Group(children),
	)
}