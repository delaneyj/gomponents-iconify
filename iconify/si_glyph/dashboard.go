package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dashboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M8.521 15.979c-4.111 0-7.459-3.36-7.459-7.489C1.062 4.359 4.41 1 8.521 1s7.458 3.359 7.458 7.49c0 4.129-3.346 7.489-7.458 7.489zM8.512 2.956c-3.052 0-5.534 2.486-5.534 5.545c0 3.058 2.482 5.545 5.534 5.545c3.051 0 5.532-2.487 5.532-5.545c0-3.059-2.482-5.545-5.532-5.545z"/><path d="M8.965 10.062V7.639c0-.347-.213-.626-.477-.626s-.479.279-.479.626v2.423c-.387.192-.968.542-.968 1.438c0 .896.819 1.484 1.446 1.484s1.428-.589 1.428-1.484s-.56-1.245-.95-1.438zM8 4h.959v.976H8zM5 5h.998v.979H5zm6 0h.998v.979H11z"/></g>`),
		g.Group(children),
	)
}