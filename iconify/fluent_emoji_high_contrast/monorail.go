package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Monorail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 7c-7.63 1.2-14.81 4.582-20.701 9.82l-6.44 5.731C1.128 24.089 2.188 27 4.478 27H2v3h28V7ZM11 19c-.823 0-1.63-.071-2.413-.207l1.377-1.225a39.814 39.814 0 0 1 14.339-8.193C22.466 14.965 17.204 19 11 19Zm4.902 6.269c1.833-2.036 4.746-3.277 7.898-3.269H29v4H15.244l.658-.731ZM9 25H7c-.55 0-1-.45-1-1s.45-1 1-1h2c.55 0 1 .45 1 1c0 .56-.45 1-1 1Z"/>`),
		g.Group(children),
	)
}