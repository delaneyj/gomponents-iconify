package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WhiteQuestionMark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M18 27.5a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0Zm-2.408-4.278c-1.41 0-2.555-1.057-2.555-2.36v-2.79c0-1.723 1.336-3.201 3.17-3.514c1.559-.265 2.683-1.528 2.683-2.996v-1.244c0-1.694-1.442-3.103-3.223-3.143a3.394 3.394 0 0 0-2.375.872c-.636.577-.986 1.35-.986 2.173c0 1.302-1.145 2.36-2.555 2.36c-1.41 0-2.566-1.058-2.566-2.37c0-2.104.901-4.082 2.523-5.55c1.633-1.47 3.785-2.262 6.065-2.213C20.31 2.535 24 6.069 24 10.308v1.254c0 3.357-2.323 6.289-5.673 7.347a.255.255 0 0 0-.18.242v1.702c0 1.312-1.144 2.37-2.555 2.37Z"/>`),
		g.Group(children),
	)
}