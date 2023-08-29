package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lavabo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M13.969 5.021H4.031c0 1.643 2.489 3.646 4.866 3.893l-.849 5.697c0 .738.548 1.386 1.286 1.386h3.299c.738 0 1.336-.598 1.336-1.336v-9.64z"/><path d="M12.062 5.027V1.523c0-.396-.219-.537-.607-.537h-1.09l-.412.967l-.911-.284l.653-1.328a.487.487 0 0 1 .448-.296h1.499c.789 0 1.299.459 1.299 1.168v3.814h-.879z"/><path d="M11 3h1.223v.937H11z"/></g>`),
		g.Group(children),
	)
}