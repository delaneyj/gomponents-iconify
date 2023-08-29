package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Androidrun(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.72 5.5a3.16 3.16 0 1 1-3.16 3.16a3.15 3.15 0 0 1 3.16-3.16Zm-12.25 6.27h8a1.1 1.1 0 0 1 .79.33l5.24 5.22l3.66-4.45a1.09 1.09 0 0 1 1.53-.16l.75.55a1.09 1.09 0 0 1 .17 1.57L36.52 21a1.09 1.09 0 0 1-1.62.07l-3.72-3.54l-4.78 5.3l4.77 4.81a1.07 1.07 0 0 1 .29 1l-2.73 13a1.09 1.09 0 0 1-1.36.84L26 42.1a1.08 1.08 0 0 1-.79-1.28l2.32-10.61l-5-5l-4.94 5.55a1.38 1.38 0 0 1-.25.21a1.85 1.85 0 0 1-.9.36H7.3a1.15 1.15 0 0 1-1.15-1.14v-1.13a1.21 1.21 0 0 1 1.15-1.15h7.47L25.68 14.7h-3.45l-4.59 5.74a1.18 1.18 0 0 1-1.54.28l-.6-.46a1.08 1.08 0 0 1-.16-1.54l5.28-6.46a1.11 1.11 0 0 1 .85-.43Z"/>`),
		g.Group(children),
	)
}