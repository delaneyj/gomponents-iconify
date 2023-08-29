package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rumble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.453 13.546a1.88 1.88 0 0 0 .275-2.645a1.88 1.88 0 0 0-.275-.275A21.213 21.213 0 0 0 10.14 7.85c-1.066-.51-2.256.2-2.426 1.414a23.523 23.523 0 0 0-.14 5.502c.116 1.23 1.292 1.964 2.372 1.492a19.628 19.628 0 0 0 4.506-2.704v-.008zm6.932-5.4a5.85 5.85 0 0 1 .014 7.872a26.149 26.149 0 0 1-13.104 7.828a5.116 5.116 0 0 1-6.167-3.578c-1.524-5.2-1.3-11.08.17-16.305C3.07 1.22 5.651-.503 8.308.131c4.925 1.174 9.545 4.196 13.077 8.013v.001z"/>`),
		g.Group(children),
	)
}