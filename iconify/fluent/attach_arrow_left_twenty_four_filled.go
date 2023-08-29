package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AttachArrowLeftTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M16 2a6 6 0 0 1 4.638 9.806a6.456 6.456 0 0 0-2.069-.718l.23-.23l.159-.165a4 4 0 0 0-5.753-5.554l-.155.16l-.018.012l-9.326 9.33a1 1 0 0 1-1.414-1.415L11.6 3.913l.046-.043A5.985 5.985 0 0 1 16 2z" fill="currentColor"/><path d="M11.087 18.568a6.46 6.46 0 0 0 .733 2.096l-.341.341l-.053.05l-.056.045a3.721 3.721 0 0 1-5.253-5.242l.149-.164l.015-.011l7.29-7.304a1 1 0 0 1 1.416 1.413l-7.29 7.304l-.012.008a1.721 1.721 0 0 0 2.289 2.553l.122-.1l.001.001l.99-.99z" fill="currentColor"/><path d="M23 17.5a5.5 5.5 0 1 1-11 0a5.5 5.5 0 0 1 11 0zm-8.5-.5a.5.5 0 0 0 0 1h4.793l-1.647 1.646a.5.5 0 0 0 .707.708l2.5-2.5a.5.5 0 0 0 0-.708l-2.5-2.5a.5.5 0 0 0-.707.708L19.293 17H14.5z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}