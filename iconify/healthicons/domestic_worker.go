package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DomesticWorker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M24 25c3.437 0 6-2.582 6-5.5S27.437 14 24 14s-6 2.582-6 5.5s2.563 5.5 6 5.5Zm0 2c4.418 0 8-3.358 8-7.5c0-4.142-3.582-7.5-8-7.5s-8 3.358-8 7.5c0 4.142 3.582 7.5 8 7.5Z"/><path d="m24 5.943l4.622 1.584A5 5 0 0 1 32 12.257V19h-2v-6.743a3 3 0 0 0-2.027-2.838L24 8.057L20.027 9.42A3 3 0 0 0 18 12.257V19h-2v-6.743a5 5 0 0 1 3.378-4.73L24 5.943ZM20 37a4 4 0 0 0 4-4a4 4 0 1 0 6.158-3.369C35.854 30.697 42 33.122 42 37v5H6v-5c0-3.878 6.146-6.303 11.842-7.369A4 4 0 0 0 20 37Zm-6 5v-6h-2v6h2Zm22-6v6h-2v-6h2Z"/><path d="M17.06 29.786C17.509 32.728 20.444 35 24 35c3.555 0 6.491-2.272 6.94-5.214a40.204 40.204 0 0 0-1.94-.354v.068c0 1.93-2.232 3.495-4.988 3.5h-.024C21.232 32.995 19 31.43 19 29.5v-.068c-.638.1-1.288.218-1.94.354Z"/></g>`),
		g.Group(children),
	)
}