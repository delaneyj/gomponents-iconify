package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BounceLeftFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M8.486 11.143a1 1 0 0 1 1.371.343c1.045 1.74 1.83 3.443 2.392 5.237l.172.581l.092-.13c2.093-2.921 4.48-3.653 7.565-2.7l.238.077a1 1 0 1 1-.632 1.898c-2.932-.978-4.73-.122-6.79 3.998c-.433.866-1.721.673-1.88-.283c-.46-2.76-1.369-5.145-2.871-7.65a1 1 0 0 1 .343-1.371zM6 4a3 3 0 1 0 0 6a3 3 0 0 0 0-6z"/></g>`),
		g.Group(children),
	)
}