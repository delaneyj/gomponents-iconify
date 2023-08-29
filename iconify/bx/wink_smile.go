package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WinkSmile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2zm0 18c-4.411 0-8-3.589-8-8s3.589-8 8-8s8 3.589 8 8s-3.589 8-8 8z"/><path fill="currentColor" d="M14.828 14.828a3.988 3.988 0 0 1-2.02 1.09a4.108 4.108 0 0 1-1.616 0a4.103 4.103 0 0 1-.749-.232a4.161 4.161 0 0 1-.679-.368a4.115 4.115 0 0 1-1.082-1.082l-1.658 1.117c.215.319.462.619.733.889a5.991 5.991 0 0 0 8.485.002c.272-.271.52-.571.734-.891l-1.658-1.117c-.143.211-.307.41-.49.592z"/><circle cx="8.5" cy="10.5" r="1.5" fill="currentColor"/><path fill="currentColor" d="M15.5 10c-2 0-2.5 2-2.5 2h5s-.501-2-2.5-2z"/>`),
		g.Group(children),
	)
}