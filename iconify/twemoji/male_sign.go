package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MaleSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#269" d="M36 32a4 4 0 0 1-4 4H4a4 4 0 0 1-4-4V4a4 4 0 0 1 4-4h28a4 4 0 0 1 4 4v28z"/><path fill="#FFF" d="M15 30.5c-4.687 0-8.5-3.813-8.5-8.5s3.813-8.5 8.5-8.5s8.5 3.813 8.5 8.5s-3.813 8.5-8.5 8.5zm0-14c-3.032 0-5.5 2.468-5.5 5.5s2.468 5.5 5.5 5.5s5.5-2.468 5.5-5.5s-2.468-5.5-5.5-5.5z"/><path fill="#FFF" d="M28.5 7h-8.969a1.5 1.5 0 1 0 0 3h5.348l-7 6.879L20 19l7-6.879V17.5a1.5 1.5 0 0 0 3 0v-9A1.5 1.5 0 0 0 28.5 7z"/>`),
		g.Group(children),
	)
}