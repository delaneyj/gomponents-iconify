package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cupcake(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m19 14l-.804 5.626c-.07.487-.104.731-.222.915a1 1 0 0 1-.426.369c-.197.09-.443.09-.933.09H14m5-7h-5m5 0c1.303-.604 2-2.236 2-3.666c0-1.536-1.03-2.85-2.49-3.397a.787.787 0 0 1-.51-.729c0-1.265-1.12-2.292-2.5-2.292c-.226 0-.445.028-.654.08c-.432.107-.915.083-1.287-.145A5.833 5.833 0 0 0 10.5 3C7.462 3 5 5.257 5 8.042c0 .352-.23.674-.557.857C3.578 9.38 3 10.254 3 11.25c0 1.277.712 2.44 2 2.75m0 0l.804 5.626v.002c.07.486.104.73.222.913a1 1 0 0 0 .426.369c.197.09.443.09.933.09H10m-5-7h5m0 0h4m-4 0v7m4-7v7m0 0h-4"/>`),
		g.Group(children),
	)
}