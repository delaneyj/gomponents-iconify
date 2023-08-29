package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MentionTwentyFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M20.226 7.25c-2.623-4.542-8.432-6.098-12.974-3.475c-4.543 2.622-6.099 8.431-3.477 12.974c2.623 4.542 8.431 6.099 12.974 3.477a.75.75 0 0 1 .75 1.299c-5.26 3.037-11.987 1.235-15.024-4.026C-.562 12.24 1.24 5.512 6.501 2.475C11.76-.562 18.488 1.24 21.525 6.501a10.959 10.959 0 0 1 1.455 4.826c.013.056.02.113.02.173v2.25a3.5 3.5 0 0 1-6.623 1.581a5.5 5.5 0 1 1 1.112-3.682a.802.802 0 0 1 .011.129v1.972a2 2 0 1 0 4 0v-1.766a9.456 9.456 0 0 0-1.274-4.733ZM16 12a4 4 0 1 0-8 0a4 4 0 0 0 8 0Z"/>`),
		g.Group(children),
	)
}