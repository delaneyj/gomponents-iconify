package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationQuestionSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3.774 8.877a8.038 8.038 0 0 1 8.01-7.377h.432a8.038 8.038 0 0 1 8.01 7.377a8.693 8.693 0 0 1-1.933 6.217L13.5 20.956a1.937 1.937 0 0 1-3 0l-4.792-5.862a8.693 8.693 0 0 1-1.934-6.217Zm6.476.694c0-1.026.803-1.821 1.75-1.821s1.75.795 1.75 1.821a1.82 1.82 0 0 1-1.168 1.719c-.589.216-1.332.786-1.332 1.71a.75.75 0 0 0 1.5 0c0-.016.004-.055.061-.118a.735.735 0 0 1 .287-.184c1.264-.463 2.152-1.696 2.152-3.127c0-1.814-1.435-3.321-3.25-3.321S8.75 7.757 8.75 9.571a.75.75 0 0 0 1.5 0ZM12 17a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}