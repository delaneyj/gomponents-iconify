package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SupabaseLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 2.598V13.97H3.9c-.67 0-1.07-.785-.643-1.336L11 2.598Zm2 5.432V2.333C13 .52 10.703-.292 9.582 1.162L1.673 11.41c-1.427 1.85-.125 4.559 2.227 4.559H11v5.698c0 1.811 2.296 2.624 3.418 1.171l7.908-10.249c1.427-1.85.126-4.559-2.227-4.559H13Zm0 2h7.1c.669 0 1.069.785.643 1.337L13 21.402V10.03Z"/>`),
		g.Group(children),
	)
}