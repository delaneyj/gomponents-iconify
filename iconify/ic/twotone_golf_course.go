package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneGolfCourse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<defs><path id="icTwotoneGolfCourse0" fill="currentColor" d="M17 5.92L9 2v18H7v-1.73c-1.79.35-3 .99-3 1.73c0 1.1 2.69 2 6 2s6-.9 6-2c0-.99-2.16-1.81-5-1.97V8.98l6-3.06z" opacity=".3"/></defs><circle cx="19.5" cy="19.5" r="1.5" fill="currentColor" opacity=".3"/><use href="#icTwotoneGolfCourse0" opacity=".3"/><circle cx="19.5" cy="19.5" r="1.5" fill="currentColor"/><use href="#icTwotoneGolfCourse0"/>`),
		g.Group(children),
	)
}