package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WatchEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M9.077 5.154h-.112A3.483 3.483 0 0 0 7 2.35V1H4v1.35c-1.18.563-2 1.756-2 3.15s.82 2.587 2 3.15V10h3V8.65a3.482 3.482 0 0 0 1.965-2.804h.112a.346.346 0 1 0 0-.692zM5.5 8a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5z" fill="currentColor"/><path d="M6.5 5H6V4a.5.5 0 1 0-1 0v1.5a.5.5 0 0 0 .5.5h1a.5.5 0 0 0 0-1z" fill="currentColor"/>`),
		g.Group(children),
	)
}