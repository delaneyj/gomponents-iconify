package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func History(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M8.479 1.046c3.859 0 6.819 3.192 7.166 6.985h1.324l-1.892 1.952l-2.065-1.952h1.338c-.33-3.229-2.746-5.958-5.936-5.958c-3.412 0-6.189 2.888-6.189 6.437c0 3.549 2.777 6.438 6.189 6.438c1.695 0 3.231-.713 4.35-1.865l.822.826a7.369 7.369 0 0 1-5.107 2.065c-4.092 0-7.419-3.349-7.419-7.464s3.327-7.464 7.419-7.464z"/><path d="M8.058 4.03L8 8.953L12 9V8l-3.032.062V4.031h-.91z"/></g>`),
		g.Group(children),
	)
}