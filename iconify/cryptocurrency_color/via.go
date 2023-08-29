package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Via(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><circle cx="16" cy="16" r="16" fill="#565656"/><path fill="#fff" fill-rule="nonzero" d="M11.133 14.296H8.005v-1.719h2.47L8.58 7.627L10.144 7l3.55 9.267l4.601.03L21.856 7l1.565.627l-1.896 4.95h2.47v1.72h-3.128l-.771 2.01l3.904.025l-.01 1.719l-4.55-.029L16 27l-3.456-9.021L8 17.949l.01-1.718l3.874.025zm3.22 3.694L16 22.288l1.638-4.277z"/></g>`),
		g.Group(children),
	)
}