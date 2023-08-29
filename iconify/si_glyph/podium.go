package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Podium(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13.432 14.035h-.479V5.72a.69.69 0 0 0-.688-.692L11 5.254V3.671l.068.027a.651.651 0 0 0 .854-.351a.646.646 0 0 0-.354-.847L8.984 1.354A1.472 1.472 0 0 0 7.515.033c-.819 0-1.484.658-1.484 1.469s.665 1.469 1.484 1.469c.392 0 .745-.153 1.011-.398l1.48.691v2.175l-5.272.976a.69.69 0 0 0-.689.693v6.93h-.539c-.279 0-.506.429-.506.96c0 .532.227.963.506.963h9.926c.279 0 .506-.431.506-.963c0-.534-.227-.963-.506-.963z"/>`),
		g.Group(children),
	)
}