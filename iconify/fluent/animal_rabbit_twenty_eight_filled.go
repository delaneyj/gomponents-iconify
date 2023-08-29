package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AnimalRabbitTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m18.184 5.55l6.555 6.555a4.305 4.305 0 0 1-4.038 7.233v.247a2.415 2.415 0 0 1-2.415 2.416H15.5v-.826A3.173 3.173 0 0 0 12.335 18h-1.59a.75.75 0 0 0 0 1.5h1.59c.917 0 1.666.75 1.666 1.674V22H7.399A2.401 2.401 0 0 1 5 19.585v-3.15c0-.233.015-.463.043-.69A2.875 2.875 0 1 1 7.55 11.82a5.397 5.397 0 0 1 2.868-.82h4.849a5.43 5.43 0 0 1 2.06.404c.197-.332.438-.615.72-.905l-2.406-2.406a1.798 1.798 0 1 1 2.543-2.543Z"/>`),
		g.Group(children),
	)
}