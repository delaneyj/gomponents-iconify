package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TranslationBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M22 12a9.969 9.969 0 0 1-2.922 7.064M19.141 5a10.015 10.015 0 0 1 2.027 3M5 19.142A9.97 9.97 0 0 1 2 12a9.969 9.969 0 0 1 2.936-7.078M6 11.982C6 10.427 6.673 9.018 7.762 8m8.523.044C17.345 9.059 18 10.449 18 11.982c0 1.552-.67 2.957-1.753 3.974M7.8 16a5.738 5.738 0 0 1-.865-1"/><circle cx="12" cy="12" r="2"/></g>`),
		g.Group(children),
	)
}