package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckBoxWithCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M1 5.125A4.125 4.125 0 0 1 5.125 1h21.75A4.125 4.125 0 0 1 31 5.125v21.75A4.125 4.125 0 0 1 26.875 31H5.125A4.125 4.125 0 0 1 1 26.875V5.125Zm11.183 17.444c.293.288.676.431 1.059.431c.383 0 .767-.143 1.059-.43l11.26-11.06a1.452 1.452 0 0 0 0-2.08a1.517 1.517 0 0 0-2.118 0L13.242 19.45l-4.685-4.602a1.517 1.517 0 0 0-2.118 0a1.453 1.453 0 0 0 0 2.08l5.744 5.641Z"/>`),
		g.Group(children),
	)
}