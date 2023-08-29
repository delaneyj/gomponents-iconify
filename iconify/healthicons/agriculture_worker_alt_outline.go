package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AgricultureWorkerAltOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M19.5 14a4 4 0 1 1 0-8a4 4 0 0 1 0 8Zm0-2a2 2 0 1 1 0-4a2 2 0 0 1 0 4Z" clip-rule="evenodd"/><path d="M13 22a2 2 0 0 1 2-2v4a2 2 0 0 1-2-2Z"/><path fill-rule="evenodd" d="M33 21.236A2.989 2.989 0 0 1 31 22h-5v17a3 3 0 1 1-6 0v-8h-1v8c0 .701-.24 1.346-.644 1.857A3 3 0 0 1 13 39V27.718c-1.563-.72-4-2.808-4-6.147C9 17.364 12.871 16 13.985 16H31c.768 0 1.47.289 2 .764V15h-1a3 3 0 0 1-3-3V6h2v6a1 1 0 0 0 1 1h1V6h2v7h1a1 1 0 0 0 1-1V6h2v6a3 3 0 0 1-3 3h-1v26a1 1 0 1 1-2 0V21.236Zm-19.163 4.665l1.163.536V39a1 1 0 1 0 2 0V29h5v10a1 1 0 1 0 2 0V20h7a1 1 0 1 0 0-2H13.985c.008 0 .006 0-.008.003c-.049.007-.24.037-.543.15a4.28 4.28 0 0 0-1.126.63c-.703.55-1.308 1.394-1.308 2.788c0 2.283 1.693 3.803 2.837 4.33Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}