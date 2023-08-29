package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func House(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><g fill="currentColor" clip-path="url(#fluentEmojiHighContrastHouse0)"><path d="M23.14 21.002h-4.28c-.48 0-.86-.38-.86-.86v-4.28c0-.47.38-.86.86-.86h4.28c.47 0 .86.38.86.86v4.28c0 .48-.38.86-.86.86Z"/><path d="m18.28.923l.004.005l12.755 12.565l.003.003a3.172 3.172 0 0 1-.003 4.546a3.192 3.192 0 0 1-2.039.916v6.151a3.91 3.91 0 0 1 3 3.803v2.09H0v-2.09a3.904 3.904 0 0 1 3-3.804v-6.11a3.232 3.232 0 0 1-2.04-.917a3.183 3.183 0 0 1-.002-4.555l.003-.002L4 10.532v-7.01C4 2.059 5.208 1 6.543 1h2.924c1.102 0 2.092.72 2.42 1.769L13.752.93c1.26-1.252 3.28-1.228 4.526-.008ZM10 7.432v-3.91A.528.528 0 0 0 9.467 3H6.543A.53.53 0 0 0 6 3.523v7.846l4-3.937Zm-5 8.314v11.256h2c0-.55.45-1 1-1v-9.61c0-.75.61-1.36 1.36-1.36h5.29c.75 0 1.36.61 1.36 1.36v9.612c.527.026.95.465.95.998H27V15.706L16.02 4.893L5 15.745Zm10 4.756a.5.5 0 1 0-1 0a.5.5 0 0 0 1 0Z"/></g><defs><clipPath id="fluentEmojiHighContrastHouse0"><path fill="#fff" d="M0 0h32v32H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}