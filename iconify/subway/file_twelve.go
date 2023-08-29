package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTwelve(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M362.3 0v128h128L362.3 0zM63.7 85.3H21V512h384v-42.7H63.7v-384zM341 0H106.3v426.7h384V149.3H341V0zm64 277.3L298.3 384L191.7 277.3h64V192H341v85.3h64z"/>`),
		g.Group(children),
	)
}