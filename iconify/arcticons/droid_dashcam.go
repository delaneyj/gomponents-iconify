package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DroidDashcam(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.71 38.911v1.867a1.73 1.73 0 0 1-1.743 1.722H7.479a1.729 1.729 0 0 1-1.743-1.722V20.67c0-.954.777-1.722 1.743-1.722h13.424M40.71 29.263v2.923m0-9.649h.98c.318 0 .574.256.574.574v5.578a.573.573 0 0 1-.574.574h-.98"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.42 18.949h13.547a1.73 1.73 0 0 1 1.743 1.722v1.866M20.77 12.204h-4.554a2.292 2.292 0 0 1-2.297-2.298V7.797A2.292 2.292 0 0 1 16.216 5.5h13.89a2.292 2.292 0 0 1 2.297 2.297v2.11a2.292 2.292 0 0 1-2.297 2.297h-4.554M10.413 22.21h25.62c.318 0 .574.256.574.575v15.879a.573.573 0 0 1-.574.574h-25.62a.573.573 0 0 1-.574-.574v-15.88c0-.318.256-.574.574-.574Zm30.297 9.975h.98c.318 0 .574.256.574.574v5.578a.573.573 0 0 1-.574.574h-.98"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.901 18.95a2.31 2.31 0 0 1-.13-.77v-8.081a2.292 2.292 0 0 1 2.296-2.297h.188a2.292 2.292 0 0 1 2.297 2.297v8.081c0 .27-.046.529-.131.77m-4.518-.001h4.516"/>`),
		g.Group(children),
	)
}