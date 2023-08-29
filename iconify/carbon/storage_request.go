package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StorageRequest(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4 21H2v3a2.006 2.006 0 0 0 2 2h3v-2H4zM4 8h3V6H4a2.006 2.006 0 0 0-2 2v3h2zm13-2h6v2h-6zM9 6h6v2H9zm8 18h6v2h-6zm11-9H4v-2H2v6h2v-2h24v2h2v-6h-2v2zm0 9h-3v2h3a2.006 2.006 0 0 0 2-2v-3h-2zm0-18h-3v2h3v3h2V8a2.006 2.006 0 0 0-2-2zM9 24h6v2H9z"/><circle cx="7" cy="12" r="1" fill="currentColor"/><circle cx="7" cy="20" r="1" fill="currentColor"/>`),
		g.Group(children),
	)
}