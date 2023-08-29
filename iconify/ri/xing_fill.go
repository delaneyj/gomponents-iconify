package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func XingFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.462 3.23c.154 0 .308.077.384.154a.49.49 0 0 1 0 .462l-6.076 10.77l3.846 7.076a.49.49 0 0 1 0 .462a.587.587 0 0 1-.385.153h-2.77c-.384 0-.614-.307-.768-.538l-3.923-7.154C11 14.307 16.924 3.77 16.924 3.77c.153-.308.384-.539.769-.539h2.769ZM8.923 7c.385 0 .616.307.77.538l1.923 3.308c-.154.154-3 5.23-3 5.23c-.154.231-.385.54-.77.54H5.155a.587.587 0 0 1-.384-.155a.49.49 0 0 1 0-.461l2.846-5.154l-1.846-3.23a.49.49 0 0 1 0-.462A.588.588 0 0 1 6.154 7h2.77Z"/>`),
		g.Group(children),
	)
}