package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Renal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M16 44c-1.462-4.349-1.934-8.016-1.415-11c.518-2.985 2.13-4.935 4.836-5.852c-2.013 6.107-1.188 10.568 2.475 13.385c5.495 4.225 18.738 2.282 19.455-12.01c.718-14.292-6.156-24.61-16.32-24.61c-10.165 0-11.685 11-4.526 14.734c-6.334 1.38-10.353 4.214-12.058 8.5C6.742 31.435 6.61 37.053 8.05 44"/><path d="M30.6 15c-2.898 1.727-3.948 3.31-3.15 4.746c.797 1.437 2.605 1.437 5.424 0m.373 6.494c-3.74-.762-5.99-.486-6.748.828c-1.137 1.97 3.87 4.826 5.686 5.183"/></g>`),
		g.Group(children),
	)
}