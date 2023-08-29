package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Imdb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.038 24.925a2.739 2.739 0 0 1 2.731-2.73h0a2.739 2.739 0 0 1 2.731 2.73V26.7a2.739 2.739 0 0 1-2.731 2.731h0a2.739 2.739 0 0 1-2.73-2.73m-.001 2.73V18.508M24.5 29.424V18.5h1.776a5.478 5.478 0 0 1 5.462 5.462h0a5.478 5.478 0 0 1-5.462 5.462Zm-13.422.007V18.508l5.462 10.923l5.462-10.923v10.923M8.5 18.508v10.923"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 37.5v-27a2 2 0 0 0-2-2h-35a2 2 0 0 0-2 2v27a2 2 0 0 0 2 2h35a2 2 0 0 0 2-2Z"/>`),
		g.Group(children),
	)
}