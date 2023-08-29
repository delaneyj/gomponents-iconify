package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsAward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M5 8.999a6.99 6.99 0 0 0 2.879 5.646l.001.001a6.972 6.972 0 0 0 1.881.979l.051.019a6.906 6.906 0 0 0 1.163.271a6.79 6.79 0 0 0 1.024.085H12c.35 0 .69-.034 1.027-.084l.182-.028c.336-.059.664-.139.981-.243l.042-.016C17 14.693 19 12.078 19 8.999C19 5.14 15.86 2 12 2S5 5.14 5 8.999zM12 4c2.756 0 5 2.242 5 4.999h-2A3.003 3.003 0 0 0 12 6V4zM7.521 16.795V22L12 20.5l4.479 1.5l.001-5.205a8.932 8.932 0 0 1-8.959 0z" fill="currentColor"/>`),
		g.Group(children),
	)
}