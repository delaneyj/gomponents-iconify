package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandHipchat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M17.802 17.292s.077-.055.2-.149c1.843-1.425 3-3.49 3-5.789c0-4.286-4.03-7.764-9-7.764s-9 3.478-9 7.764c0 4.288 4.03 7.646 9 7.646c.424 0 1.12-.028 2.088-.084c1.262.82 3.104 1.493 4.716 1.493c.499 0 .734-.41.414-.828c-.486-.596-1.156-1.551-1.416-2.29z"/><path d="M7.5 13.5c2.5 2.5 6.5 2.5 9 0"/></g>`),
		g.Group(children),
	)
}