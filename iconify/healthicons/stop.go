package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path d="M30 25.14V12.5a1.5 1.5 0 1 0-3 0V23h-2V9.5a1.5 1.5 0 1 0-3 0V23h-2V12.5a1.5 1.5 0 1 0-3 0V25h-2v-8.5a1.5 1.5 0 1 0-3 0v13.238a10.263 10.263 0 0 0 19.324 4.818l5.906-11.107a1.812 1.812 0 0 0-3.184-1.731l-2.172 3.909c-.5.902-1.874.546-1.874-.486Z"/><path fill-rule="evenodd" d="M24 4c11.046 0 20 8.954 20 20s-8.954 20-20 20S4 35.046 4 24S12.954 4 24 4Zm18.095 20c0-9.994-8.101-18.095-18.095-18.095c-9.994 0-18.095 8.101-18.095 18.095c0 9.994 8.101 18.095 18.095 18.095c9.994 0 18.095-8.101 18.095-18.095Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}