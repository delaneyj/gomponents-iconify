package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EuroCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopEuroCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M15.489 7C12.43 7 10 9.213 10 12.5c0 3.386 2.527 6 5.489 6c.743 0 1.451-.161 2.098-.454a1 1 0 1 1 .826 1.821a7.061 7.061 0 0 1-2.924.633C11.283 20.5 8 16.846 8 12.5C8 8.055 11.38 5 15.489 5c1.237 0 2.428.393 3.574 1.174a1 1 0 1 1-1.126 1.652C17.079 7.242 16.274 7 15.489 7Z"/><path d="M6 11a1 1 0 0 1 1-1h9a1 1 0 1 1 0 2H7a1 1 0 0 1-1-1Zm0 3.5a1 1 0 0 1 1-1h8a1 1 0 1 1 0 2H7a1 1 0 0 1-1-1Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopEuroCircleFilled0)"/></g>`),
		g.Group(children),
	)
}