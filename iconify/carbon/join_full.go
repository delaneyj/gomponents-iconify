package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JoinFull(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" d="M16 9.075a7.994 7.994 0 0 1 0 13.85a7.994 7.994 0 0 1 0-13.85m0-2.237a9.995 9.995 0 0 0 0 18.324a9.995 9.995 0 0 0 0-18.324Z"/><path fill="currentColor" d="M10 16a9.998 9.998 0 0 1 6-9.162a10 10 0 1 0 0 18.324A9.998 9.998 0 0 1 10 16Z"/><path fill="currentColor" d="M16 9.075a7.994 7.994 0 0 0 0 13.85a7.994 7.994 0 0 0 0-13.85Z"/><path fill="currentColor" d="M20 6a9.954 9.954 0 0 0-4 .838a9.995 9.995 0 0 1 0 18.324A9.999 9.999 0 1 0 20 6Z"/>`),
		g.Group(children),
	)
}