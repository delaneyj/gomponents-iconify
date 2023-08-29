package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserVoiceFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 22a8 8 0 0 1 16 0H1Zm8-9c-3.315 0-6-2.685-6-6s2.685-6 6-6s6 2.685 6 6s-2.685 6-6 6Zm9.246-9.815A9.97 9.97 0 0 1 19 7a9.97 9.97 0 0 1-.754 3.816l-1.677-1.22A7.99 7.99 0 0 0 17 7a7.99 7.99 0 0 0-.43-2.596l1.676-1.22ZM21.548.784A13.942 13.942 0 0 1 23 7c0 2.233-.523 4.344-1.452 6.217l-1.645-1.197A11.956 11.956 0 0 0 21 7a11.96 11.96 0 0 0-1.097-5.02L21.548.784Z"/>`),
		g.Group(children),
	)
}