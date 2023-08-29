package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RwOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#20603d" d="M0 0h512v512H0z"/><path fill="#fad201" d="M0 0h512v384H0z"/><path fill="#00a1de" d="M0 0h512v256H0z"/><g transform="translate(374.4 133.8) scale(.7111)"><g id="flagRw1x10"><path id="flagRw1x11" fill="#e5be01" d="M116.1 0L35.7 4.7l76.4 25.4l-78.8-16.3L100.6 58l-72-36.2L82 82.1L21.9 28.6l36.2 72l-44.3-67.3L30 112L4.7 35.7L0 116.1L-1-1z"/><use width="100%" height="100%" href="#flagRw1x11" transform="scale(1 -1)"/></g><use width="100%" height="100%" href="#flagRw1x10" transform="scale(-1 1)"/><circle r="34.3" fill="#e5be01" stroke="#00a1de" stroke-width="3.4"/></g>`),
		g.Group(children),
	)
}