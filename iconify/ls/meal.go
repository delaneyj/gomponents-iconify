package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Meal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M201 248V14h67v268c0 37-30 68-67 68v296c0 18-15 34-34 34h-67c-18 0-33-16-33-34V350c-37 0-67-31-67-68V14h67v234h33V14h67v234h34zm168 102V114c0-55 45-100 100-100s100 45 100 100v532c0 18-14 34-32 34h-68c-18 0-34-16-34-34V350h-66z"/>`),
		g.Group(children),
	)
}