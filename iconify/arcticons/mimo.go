package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mimo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.18 4.684a58.598 58.598 0 0 1 14.466 4.809c3.16 1.626 4.18 4.489 3.462 7.969a31.072 31.072 0 0 1-2.522 8.061a8.014 8.014 0 0 1-8.558 4.346a38.465 38.465 0 0 1-14.834-5.619C8.27 22.229 7.6 19.597 8.395 16.134a24.807 24.807 0 0 1 2.448-7.325c2.01-3.37 4.86-4.861 8.337-4.125Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m23.957 19.368l1.156-2.776a1.619 1.619 0 1 1 2.989 1.245l-1.156 2.775a1.619 1.619 0 0 1-2.989-1.244Zm-9.401-4.033a1.989 1.989 0 0 1 3.68 1.485"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.978 19.16a15.321 15.321 0 1 1-26.981 4.36"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.133 29.688s2.063 5.23 5.304 3.094"/>`),
		g.Group(children),
	)
}