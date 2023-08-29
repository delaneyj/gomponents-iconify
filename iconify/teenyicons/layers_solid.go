package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayersSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7.697 1.04a.5.5 0 0 0-.394 0l-7 3a.5.5 0 0 0 0 .92l7 3a.5.5 0 0 0 .394 0l7-3a.5.5 0 0 0 0-.92l-7-3Z"/><path fill="currentColor" d="M7.5 9.956L.697 7.04l-.394.92L7.5 11.044l7.197-3.084l-.394-.92L7.5 9.956Z"/><path fill="currentColor" d="m.697 10.04l-.394.92L7.5 14.044l7.197-3.084l-.394-.92L7.5 12.956L.697 10.04Z"/>`),
		g.Group(children),
	)
}