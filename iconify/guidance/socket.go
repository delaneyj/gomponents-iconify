package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Socket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M9.5 9.5v-4m5 4v-4m-2.5 12h3.5v-.492a10 10 0 0 1 2.191-6.247L18.5 9.75V9.5h-13v.25l.809 1.01A10 10 0 0 1 8.5 17.009v.492H12Zm0 0v6c6.351 0 11.5-5.149 11.5-11.5S18.351.5 12 .5S.5 5.649.5 12c0 5.493 3.85 10.086 9 11.227"/>`),
		g.Group(children),
	)
}