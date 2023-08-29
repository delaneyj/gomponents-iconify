package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sns(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="m271 392l173 97c25-29 62-47 104-47c76 0 137 62 137 138s-61 137-137 137c-77 0-138-61-138-137c0-8 1-17 2-25l-178-99c-25 25-59 40-97 40C61 496 0 435 0 359s61-137 137-137c39 0 75 16 100 42l177-96c-2-10-4-20-4-30C410 62 471 0 548 0c76 0 137 62 137 138s-61 137-137 137c-39 0-74-15-99-42l-178 95c2 10 3 21 3 31c0 11-1 23-3 33z"/>`),
		g.Group(children),
	)
}