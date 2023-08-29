package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Health(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M21.671 6.8a5.537 5.537 0 0 1 5.737 0l2.037 1.246c2.858 1.747 6.113 2.136 9.383 2.528c1.057.127 2.115.253 3.162.426c0 0 .023 1.834-.001 2.666l-.209 7.151c-.18 6.133-2.907 11.948-7.585 16.174c-2.207 1.993-4.49 3.794-7.113 5.29a5.545 5.545 0 0 1-5.1.187c-3.029-1.44-5.711-3.266-8.16-5.477c-4.678-4.226-7.405-10.041-7.584-16.174l-.208-7.128c-.025-.845 0-2.689 0-2.689c.85-.1 1.716-.183 2.589-.266c3.853-.366 7.834-.744 11.1-2.74L21.672 6.8ZM22 22v-7h4v7h7v4h-7v7h-4v-7h-7v-4h7Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}