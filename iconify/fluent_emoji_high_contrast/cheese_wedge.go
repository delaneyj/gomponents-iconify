package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheeseWedge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><g fill="currentColor" clip-path="url(#fluentEmojiHighContrastCheeseWedge0)"><path d="M24.48 16.435a3.67 3.67 0 1 1-7.34 0a3.67 3.67 0 0 1 7.34 0Zm-11.87 5.55a2.12 2.12 0 1 0 0-4.24a2.12 2.12 0 0 0 0 4.24Zm6.35 1.14a1.14 1.14 0 1 1-2.28 0a1.14 1.14 0 0 1 2.28 0Z"/><path d="M31 8.425a3.367 3.367 0 0 0-2.19-3.156L19.184.456a4.315 4.315 0 0 0-4.747.592l-10.952 9.5A3.358 3.358 0 0 0 2 13.336v9.01a2 2 0 0 0 2 2c.065 0 .12.054.12.12c0 .065-.055.12-.12.12a2 2 0 0 0-2 2v2.01c0 2.13 1.966 3.753 4.082 3.282h.004l22.274-4.91l.01-.002a3.369 3.369 0 0 0 2.63-3.28v-8.62a2 2 0 0 0-2.13-1.996a.311.311 0 0 1-.25-.304c0-.171.14-.31.31-.31H29a2 2 0 0 0 2-2v-2.03ZM4 22.345V13.99c.353.136.747.173 1.138.09l23.35-4.99c.181-.04.353-.102.512-.186v1.55h-.07a2.311 2.311 0 0 0 0 4.62c.02 0 .04 0 .07-.01v8.62c0 .64-.45 1.19-1.07 1.33l-22.28 4.91c-.85.19-1.65-.46-1.65-1.33v-2.01a2.121 2.121 0 0 0 0-4.24Zm24.93-9.27h-.006h.006Z"/></g><defs><clipPath id="fluentEmojiHighContrastCheeseWedge0"><path fill="#fff" d="M0 0h32v32H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}