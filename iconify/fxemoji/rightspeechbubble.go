package fxemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rightspeechbubble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#E3E2DD" d="M267.206 8.765c-130.1 0-235.566 105.594-235.566 235.851c0 30.486 5.778 59.62 16.297 86.367L8.015 461.423a12.606 12.606 0 0 0-.757 2.439l-.052.138h.03c-.729 4 .302 8.081 3.214 11.183c2.909 3.101 6.996 4.409 10.915 3.924a.114.114 0 0 0 0 .027l.165-.027c.865-.12 1.719-.314 2.55-.609l124.823-29.9c34.767 20.256 75.178 31.867 118.301 31.867c130.1 0 235.566-105.594 235.566-235.85S397.305 8.765 267.206 8.765z"/>`),
		g.Group(children),
	)
}