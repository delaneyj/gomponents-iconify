package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IceAccretion(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4 2a2 2 0 0 0-2 2v16l1 2l1-2V4h4v10l1 2l1-2V2zm24 0H14v8l1 2l1-2V4h2v13l1 2l1-2V4h4v10l1 2l1-2V4h2v20l1 2l1-2V4a2 2 0 0 0-2-2zM14 28l-1 2l-1-2V16h2v12z"/><path fill="currentColor" d="m24 26l-1 2l-1-2v-6h2v6zM8 24l-1 2l-1-2v-6h2v6z"/>`),
		g.Group(children),
	)
}