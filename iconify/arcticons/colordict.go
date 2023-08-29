package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Colordict(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.4 6.5v35a2 2 0 0 0 2 2h2.33v-39H10.4a2 2 0 0 0-2 2Zm4.331-2v39h24.87a2 2 0 0 0 2-2v-35a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.666 28.665v-8.838h1.988a3.866 3.866 0 0 1 3.867 3.867v1.104a3.866 3.866 0 0 1-3.867 3.867Z"/><circle cx="24.742" cy="20.104" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.742 22.81v5.855m9.488-7.673v6.573a1.105 1.105 0 0 0 1.105 1.105h.331m-2.596-5.855h2.32m-4.3 4.742a2.208 2.208 0 0 1-1.918 1.113h0a2.21 2.21 0 0 1-2.209-2.21v-1.436a2.21 2.21 0 0 1 2.21-2.209h0a2.208 2.208 0 0 1 1.916 1.11"/>`),
		g.Group(children),
	)
}