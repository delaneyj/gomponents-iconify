package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ophiuchus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M10 7a1 1 0 0 1 1 1v4.363c.683.06 1.366.211 2.031.46l6.64 2.49c.436.164.882.271 1.329.324V8a1 1 0 1 1 2 0v7.516a5.666 5.666 0 0 0 3.2-2.107a1 1 0 0 1 1.6 1.202a7.667 7.667 0 0 1-4.8 2.947V18c0 3.862-3.138 7-7 7s-7-3.138-7-7v-3.506a5.666 5.666 0 0 0-3.2 2.107A1 1 0 1 1 4.2 15.4A7.667 7.667 0 0 1 9 12.451V8a1 1 0 0 1 1-1Zm1 7.373V18c0 2.758 2.242 5 5 5s5-2.242 5-5v-.353a7.658 7.658 0 0 1-2.031-.46l-6.64-2.49A5.665 5.665 0 0 0 11 14.372Z"/><path d="M1 6a5 5 0 0 1 5-5h20a5 5 0 0 1 5 5v20a5 5 0 0 1-5 5H6a5 5 0 0 1-5-5V6Zm5-3a3 3 0 0 0-3 3v20a3 3 0 0 0 3 3h20a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3H6Z"/></g>`),
		g.Group(children),
	)
}