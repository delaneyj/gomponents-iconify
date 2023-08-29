package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudWithLightning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M2 13.5c0 .05 0 .098.002.148a5.965 5.965 0 0 0-.002.14c0 2.476 1.54 4.613 3.767 5.611A6.47 6.47 0 0 0 8.5 20h2.632l4.303-6.456c.768-1.152 2.565-.608 2.565.777V19h.787c.665 0 1.174.443 1.342 1H21c4.694 0 9-4 9-9c0-.658-.085-1.275-.244-1.849C28.915 5.07 25.3 2 20.97 2a8.945 8.945 0 0 0-6.513 2.802A4.042 4.042 0 0 0 8.133 7.01A6.5 6.5 0 0 0 2 13.5Z"/><path d="M17.244 19.968A.396.396 0 0 0 17.4 20h1.387a.4.4 0 0 1 .341.609l-4.387 7.178c-.21.345-.741.195-.741-.209V22.4a.4.4 0 0 0-.4-.4h-1.853a.4.4 0 0 1-.332-.622l4.852-7.279a.389.389 0 0 1 .383-.175c.184.024.35.172.35.397V19.6a.4.4 0 0 0 .244.368Z"/></g>`),
		g.Group(children),
	)
}