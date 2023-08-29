package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SandWatchSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m10.664 11.839l.187.161l-.187.161A21.075 21.075 0 0 0 5.82 18.22a1.428 1.428 0 0 0 1.135 2.093l1.174.102c2.576.226 5.166.226 7.742 0l1.174-.102a1.428 1.428 0 0 0 1.135-2.093a21.074 21.074 0 0 0-4.844-6.058l-.187-.16l.187-.162a21.074 21.074 0 0 0 4.844-6.057a1.428 1.428 0 0 0-1.135-2.094l-1.174-.102a44.446 44.446 0 0 0-7.742 0l-1.174.102A1.428 1.428 0 0 0 5.82 5.782a21.075 21.075 0 0 0 4.844 6.057ZM12 9.75a.75.75 0 0 1-.53-.22l-2-2A.75.75 0 0 1 10 6.25h4a.75.75 0 0 1 .53 1.28l-2 2a.75.75 0 0 1-.53.22Zm0 4.5a.75.75 0 0 0-.53.22l-2 2a.75.75 0 0 0 .53 1.28h4a.75.75 0 0 0 .53-1.28l-2-2a.75.75 0 0 0-.53-.22Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}