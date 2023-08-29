package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zzz(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M27.662 8.552a1.5 1.5 0 0 1 .159 1.66l-5.446 10.113l5.8-1.29a1.5 1.5 0 0 1 .65 2.93l-9 2a1.5 1.5 0 0 1-1.646-2.176l5.283-9.812l-6.531 1.96a1.5 1.5 0 1 1-.862-2.874l10-3a1.5 1.5 0 0 1 1.593.489Z"/><path d="M9 19a1 1 0 0 1 1-1h7a1 1 0 0 1 .753 1.659L12.203 26H17a1 1 0 1 1 0 2h-7a1 1 0 0 1-.753-1.659L14.797 20H10a1 1 0 0 1-1-1Z"/><path d="M4.287 21.042a1 1 0 0 0-.574 1.916l3.102.93L2.4 27.2a1 1 0 0 0 .404 1.78l5 1a1 1 0 0 0 .392-1.96l-2.807-.562L9.6 24.3a1 1 0 0 0-.313-1.758l-5-1.5Z"/></g>`),
		g.Group(children),
	)
}