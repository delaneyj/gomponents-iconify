package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Heart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M14.54 20.019c-1.688-2.075-3.272-2.521-4.754-1.337c-2.223 1.776-2.861 7.968-1.073 13.427c1.789 5.459 5.267 12.893 12.289 12.893c7.021 0 8.682-7.48 11.546-12.002c2.865-4.522 4.38-8.885 1.573-14.318"/><path stroke-linecap="round" d="M11 18.037A643.145 643.145 0 0 0 7 12c-1.446-2.145 2.251-4.918 4-3.032c1.166 1.258 2.715 3.11 4.647 5.557"/><path stroke-linecap="round" stroke-linejoin="round" d="M15.024 25.64c-.485-6.064-.09-10.012 1.182-11.845c1.91-2.75 5.457-3.792 8.797-3.792c1.99 0 3.806.847 5.45 2.541"/><path d="M41 12.613c.586 2.036-.37 3.897-3.316 4.318c-2.945.421-5.153 1.902-6.745 3.148c-1.593 1.246-4.44 5.026-5.003 6.923c-.562 1.898-3.776.153-4.639-.605c-.863-.757-1.712-2.416 0-4.151c1.712-1.735 1.341-2.081 1.341-3.84c0-1.76 9.362-7.58 14.635-8.112c1.172-.068 3.142.282 3.727 2.319Z" clip-rule="evenodd"/><path stroke-linecap="round" d="M23.008 4v5.263m-2.701 1.455c-3.018-3.793-5.451-5.91-7.3-6.35m3.997 2.661l.99-4.067m17.619 7.755c-.322 1.105-.322 2.14 0 3.107c.322.967 1.013 2.002 2.071 3.107"/></g>`),
		g.Group(children),
	)
}