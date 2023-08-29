package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tobojapanese(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><circle cx="23.998" cy="35.502" r="1.316" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.258 26.991a9.25 9.25 0 1 1-.15-17.164m7.784-.032a9.25 9.25 0 1 1-.165 17.159m-7.886-8.403l4.239.032m-6.184 2.486l4.239.032m-3.92-4.781l4.239.032m-.526-1.801l4.239.032m-2.422 2.454a2.52 2.52 0 0 1 3.761-.35l-1.594 2.358s-1.57 2.148-3.474.797"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.912 19.045a5.506 5.506 0 0 1 3.06 3.378m1.689-6.629a2.498 2.498 0 0 1-3.155.892"/><circle cx="28.335" cy="14.726" r=".75" fill="currentColor"/><circle cx="23.968" cy="18.407" r="9.25" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><rect width="8" height="8" x="20.039" y="31.543" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/>`),
		g.Group(children),
	)
}