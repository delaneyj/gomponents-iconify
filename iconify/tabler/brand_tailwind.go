package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandTailwind(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.667 6C9.177 6 7.623 7.222 7 9.667c.933-1.223 2.023-1.68 3.267-1.375c.71.174 1.217.68 1.778 1.24c.916.912 2 1.968 4.288 1.968c2.49 0 4.044-1.222 4.667-3.667c-.933 1.223-2.023 1.68-3.267 1.375c-.71-.174-1.217-.68-1.778-1.24C15.039 7.056 13.98 6 11.667 6zm-4 6.5c-2.49 0-4.044 1.222-4.667 3.667c.933-1.223 2.023-1.68 3.267-1.375c.71.174 1.217.68 1.778 1.24c.916.912 1.975 1.968 4.288 1.968c2.49 0 4.044-1.222 4.667-3.667c-.933 1.223-2.023 1.68-3.267 1.375c-.71-.174-1.217-.68-1.778-1.24c-.916-.912-1.975-1.968-4.288-1.968z"/>`),
		g.Group(children),
	)
}