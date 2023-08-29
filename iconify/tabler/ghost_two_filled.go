package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GhostTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="m12 1.999l.041.002l.208.003a8 8 0 0 1 7.747 7.747l.003.248l.177.006a3 3 0 0 1 2.819 2.819L23 13a3 3 0 0 1-3 3l-.001 1.696l1.833 2.75a1 1 0 0 1-.72 1.548L21 22H11c-3.445.002-6.327-2.49-6.901-5.824l-.028-.178l-.071.001a3 3 0 0 1-2.995-2.824L1 13a3 3 0 0 1 3-3l.004-.25A8 8 0 0 1 12 2zM12 12a2 2 0 0 0-2 2a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1a2 2 0 0 0-2-2zm-1.99-4l-.127.007A1 1 0 0 0 10 10l.127-.007A1 1 0 0 0 10.01 8zm4 0l-.127.007A1 1 0 0 0 14 10l.127-.007A1 1 0 0 0 14.01 8z"/></g>`),
		g.Group(children),
	)
}