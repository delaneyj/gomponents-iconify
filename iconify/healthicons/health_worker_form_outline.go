package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HealthWorkerFormOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path d="M16 20a1 1 0 1 0 0 2h7a1 1 0 1 0 0-2h-7Zm-1-4a1 1 0 0 1 1-1h15.5a1 1 0 1 1 0 2H16a1 1 0 0 1-1-1Z"/><path fill-rule="evenodd" d="M15 28a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1v-4Zm2 3v-2h2v2h-2Zm11 0a3 3 0 1 0 0-6a3 3 0 1 0 0 6Zm0-2a1 1 0 1 0 .002-1.998A1 1 0 0 0 28 29Zm-6 6.182C22 33.066 25.997 32 28 32s6 1.066 6 3.182V39H22v-3.818Zm2.055.04a.927.927 0 0 0-.055.057V37h8v-1.72a.927.927 0 0 0-.055-.059c-.164-.16-.48-.372-.976-.583c-1-.425-2.234-.638-2.969-.638c-.735 0-1.969.213-2.969.638c-.496.21-.812.423-.976.583Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M17 7a3 3 0 0 1 3-3h8a3 3 0 0 1 3 3h4a3 3 0 0 1 3 3v31a3 3 0 0 1-3 3H13a3 3 0 0 1-3-3V10a3 3 0 0 1 3-3h4Zm11 5a3 3 0 0 0 3-3h4a1 1 0 0 1 1 1v31a1 1 0 0 1-1 1H13a1 1 0 0 1-1-1V10a1 1 0 0 1 1-1h4a3 3 0 0 0 3 3h8Zm-8-6a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1h-8Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}