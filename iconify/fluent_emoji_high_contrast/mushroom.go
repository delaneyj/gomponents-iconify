package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mushroom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M17.68 14.3a2.74 2.74 0 1 1-5.48 0a2.74 2.74 0 0 1 5.48 0Z"/><path d="M12.79 4.434A12.011 12.011 0 0 1 16 4c1.62 0 3.163.32 4.573.902c3.998 1.65 6.91 5.397 7.365 9.871A12.139 12.139 0 0 1 28 16v2a2 2 0 0 1-2 2h-7.368l2.2 5.224A2 2 0 0 1 18.987 28h-6.12a2 2 0 0 1-1.876-2.691L12.947 20H6a2 2 0 0 1-2-2v-2c0-2.271.63-4.395 1.727-6.206a12.031 12.031 0 0 1 7.062-5.36ZM26 18v-2c0-.25-.01-.496-.027-.74a6.632 6.632 0 0 1-6.543-6.63c0-.668.098-1.312.28-1.92A9.973 9.973 0 0 0 16 6c-1.15 0-2.254.194-3.282.551a4.992 4.992 0 0 1-4.935 3.748A9.954 9.954 0 0 0 6 16v2h20Z"/></g>`),
		g.Group(children),
	)
}