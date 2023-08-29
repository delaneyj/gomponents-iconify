package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tiki(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M8.5 23.585h7.737M12.368 35.01V23.585m8.358 11.421V23.59M35.83 35.006V23.59m-10.619 0v11.416m6.135 0l-4.694-5.708l4.694-5.669m-4.7 5.669h-1.435"/><circle cx="20.726" cy="16.659" r="3.669"/><path d="M22.695 16.66a1.97 1.97 0 0 1-3.94 0h3.94Z"/><circle cx="35.831" cy="16.659" r="3.669"/><path d="M37.8 16.66a1.97 1.97 0 0 1-3.94 0h3.94Z"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}