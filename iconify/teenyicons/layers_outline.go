package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayersOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m7.5 1.5l.197-.46a.5.5 0 0 0-.394 0l.197.46Zm-7 3l-.197-.46a.5.5 0 0 0 0 .92L.5 4.5Zm7 3l-.197.46a.5.5 0 0 0 .394 0L7.5 7.5Zm7-3l.197.46a.5.5 0 0 0 0-.92l-.197.46Zm-7 6l-.197.46l.197.084l.197-.084l-.197-.46Zm0 3l-.197.46l.197.084l.197-.084l-.197-.46ZM7.303 1.04l-7 3l.394.92l7-3l-.394-.92Zm-7 3.92l7 3l.394-.92l-7-3l-.394.92Zm7.394 3l7-3l-.394-.92l-7 3l.394.92Zm7-3.92l-7-3l-.394.92l7 3l.394-.92ZM.303 7.96l7 3l.394-.92l-7-3l-.394.92Zm7.394 3l7-3l-.394-.92l-7 3l.394.92Zm-7.394 0l7 3l.394-.92l-7-3l-.394.92Zm7.394 3l7-3l-.394-.92l-7 3l.394.92Z"/>`),
		g.Group(children),
	)
}