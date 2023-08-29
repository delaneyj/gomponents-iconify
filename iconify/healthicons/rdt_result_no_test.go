package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RdtResultNoTest(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M40 24a6 6 0 0 1-6 6h-1.434l-4.803-5.055A2 2 0 0 0 26 22h-1.035l-3.8-4H34a6 6 0 0 1 6 6Zm-4 0a1.5 1.5 0 1 0-3 0a1.5 1.5 0 0 0 3 0Zm-5.248-.75a.751.751 0 1 1 .002 1.502a.751.751 0 0 1-.002-1.502Z" clip-rule="evenodd"/><path d="M12 24a2 2 0 0 1 2-2h2.688l-3.716-3.912A6.002 6.002 0 0 0 14 30h10.29l-3.801-4H14a2 2 0 0 1-2-2Z"/><path fill-rule="evenodd" d="M44 24c0 11.046-8.954 20-20 20S4 35.046 4 24S12.954 4 24 4s20 8.954 20 20Zm-2 0c0 5.108-2.128 9.72-5.546 12.996L11.657 10.898A17.938 17.938 0 0 1 24 6c9.941 0 18 8.059 18 18ZM24 42c4.113 0 7.903-1.38 10.934-3.7L10.278 12.35A17.929 17.929 0 0 0 6 24c0 9.941 8.059 18 18 18Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}