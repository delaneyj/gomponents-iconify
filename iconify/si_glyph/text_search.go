package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextSearch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M0 0h13.931v.983H0zm0 2h13.931v.942H0zm0 12h10.958v.951H0zm8.49-9.946C6.01 4.054 4 6.047 4 8.506c0 2.459 2.01 4.452 4.49 4.452c2.48 0 4.489-1.993 4.489-4.452c0-2.459-2.008-4.452-4.489-4.452zm0 7.964a3.54 3.54 0 1 1 0-7.08a3.54 3.54 0 1 1 0 7.08zm7.448 2.593l-1.361 1.361l-2.996-2.996s.57-.073.931-.434c.361-.362.431-.928.431-.928l2.995 2.997z"/><path d="M8.677 6.43c.526 0 .329-.4-.403-.4a2.267 2.267 0 0 0-2.279 2.256c0 .725.404.921.404.4C6.398 7.44 7.418 6.43 8.677 6.43zM0 4h3.973v.962H0zm0 2h3v.973H0zm0 2h2.98v.993H0zm0 2h3.02v.973H0zm0 2h4v.931H0z"/></g>`),
		g.Group(children),
	)
}