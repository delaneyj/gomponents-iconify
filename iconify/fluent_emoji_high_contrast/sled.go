package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M9 11a1 1 0 0 0-1 1v7H6c-.575 0-.94-.14-1.18-.3a1.635 1.635 0 0 1-.55-.647A2.69 2.69 0 0 1 4 17v-.007A1 1 0 0 0 3 16c-1 0-1 1.001-1 1.001v.024a2.38 2.38 0 0 0 .008.175a4.685 4.685 0 0 0 .472 1.747a3.632 3.632 0 0 0 1.23 1.416C4.316 20.766 5.076 21 6 21h2v.5A1.5 1.5 0 0 0 9.5 23H11v4H6c-.575 0-.94-.14-1.18-.3a1.635 1.635 0 0 1-.55-.647A2.69 2.69 0 0 1 4 25v-.007A1 1 0 0 0 3 24c-.26 0-1 0-1 1.001v.024a2.38 2.38 0 0 0 .008.175a4.685 4.685 0 0 0 .472 1.747a3.632 3.632 0 0 0 1.23 1.416C4.316 28.766 5.076 29 6 29h23a1 1 0 1 0 0-2h-3v-4h1.5a1.5 1.5 0 0 0 1.5-1.5V12a1 1 0 0 0-1-1H9Zm0 2.5h19V15H9v-1.5ZM9 17h19v1.5H9V17Zm15 6v4H13v-4h11Z"/>`),
		g.Group(children),
	)
}