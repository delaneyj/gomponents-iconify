package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaxingCrescentMoon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M22.753 9.013a2.5 2.5 0 0 1 .936 4.686a3 3 0 0 1-.936-4.686Z"/><path d="M16.014 1c8.285 0 15 6.716 15 15c0 8.284-6.715 15-15 15c-8.284 0-15-6.716-15-15c0-8.284 6.716-15 15-15Zm2.363 2.214a18.032 18.032 0 0 1 3.934 5.793a2.5 2.5 0 1 0 1.276 4.745c.094.743.142 1.5.142 2.27c0 4.984-2.026 9.495-5.299 12.754a12.95 12.95 0 0 0 5.902-2.785a3.5 3.5 0 0 1 3.67-4.955A12.959 12.959 0 0 0 29.015 16c0-6.373-4.585-11.675-10.637-12.786ZM5.49 23.634a13.05 13.05 0 0 0 5.228 4.242a3.5 3.5 0 0 0-5.228-4.242ZM15.5 7a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm-6 9a3.5 3.5 0 1 0 0-7a3.5 3.5 0 0 0 0 7Zm7 8a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Z"/></g>`),
		g.Group(children),
	)
}