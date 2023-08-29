package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RdtResultNoTestNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g fill="currentColor" clip-path="url(#healthiconsRdtResultNoTestNegative0)"><path d="M34.5 22.5a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3ZM31.504 24a.751.751 0 1 0-1.502.002a.751.751 0 0 0 1.502-.002Z"/><path fill-rule="evenodd" d="M42 24c0 5.108-2.128 9.72-5.546 12.996L11.657 10.898A17.938 17.938 0 0 1 24 6c9.941 0 18 8.059 18 18Zm-14.237.945L32.566 30H34a6 6 0 0 0 0-12H21.164l3.8 4H26a2 2 0 0 1 1.763 2.945ZM10.278 12.35L34.934 38.3A17.921 17.921 0 0 1 24 42c-9.941 0-18-8.059-18-18c0-4.443 1.61-8.51 4.278-11.65Zm2.694 5.738L16.688 22H14a2 2 0 1 0 0 4h6.489l3.8 4H14a6 6 0 0 1-1.028-11.912Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M0 0h48v48H0V0Zm44 24c0 11.046-8.954 20-20 20S4 35.046 4 24S12.954 4 24 4s20 8.954 20 20Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsRdtResultNoTestNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}