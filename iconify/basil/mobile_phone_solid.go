package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MobilePhoneSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m18.75 11l-.006.243l.006.257l-.006.245l.006.255l-.189 8.253a2.047 2.047 0 0 1-1.826 1.989c-3.147.34-6.323.34-9.47 0a2.047 2.047 0 0 1-1.827-1.989L5.25 12l.006-.257l-.006-.243l.006-.256L5.25 11l.168-6.641l.007-.131l.01-.462A2.1 2.1 0 0 1 7.32 1.728c.087-.01.175-.018.262-.027a2.82 2.82 0 0 1 .365-.06a41.778 41.778 0 0 1 8.106 0c.124.013.246.033.364.06l.263.027a2.1 2.1 0 0 1 1.884 2.038l.011.463c.003.043.005.086.006.13l.17 6.641ZM16 17.75a.75.75 0 0 0 0-1.5H8a.75.75 0 0 0 0 1.5h8Zm-4 3.95a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}