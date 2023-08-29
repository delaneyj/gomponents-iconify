package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M12.46 8a1.5 1.5 0 0 0-1.5 1.5v14.016a1.5 1.5 0 1 0 3 0v-5.29a.25.25 0 0 1 .25-.25H18c1.379 0 2.657-.475 3.6-1.383c.948-.913 1.47-2.181 1.47-3.593c0-2.932-2.217-5-5.07-5h-5.54Zm1.75 6.977a.25.25 0 0 1-.25-.25V11.25a.25.25 0 0 1 .25-.25H18c1.276 0 2.07.803 2.07 2c0 .653-.23 1.123-.551 1.432c-.327.314-.833.545-1.519.545h-3.79Z"/><path d="M6 1a5 5 0 0 0-5 5v20a5 5 0 0 0 5 5h20a5 5 0 0 0 5-5V6a5 5 0 0 0-5-5H6ZM3 6a3 3 0 0 1 3-3h20a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V6Z"/></g>`),
		g.Group(children),
	)
}