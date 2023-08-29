package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatBubbleSolidBadged(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M30 13.25a7.46 7.46 0 0 1-4.35-1.4a1 1 0 0 1-.93.65H11.28a1 1 0 0 1 0-2H24.2A7.46 7.46 0 0 1 23 3.2a18 18 0 0 0-5-.7c-8.82 0-16 6.28-16 14s7.18 14 16 14a18 18 0 0 0 4.88-.68l5.53 3.52a1 1 0 0 0 1.54-.84v-6.73a13 13 0 0 0 4-9.27a12.34 12.34 0 0 0-.68-4a7.46 7.46 0 0 1-3.27.75Zm-8.25 9.25h-7.5a1 1 0 0 1 0-2h7.5a1 1 0 0 1 0 2Zm3.25-5H11a1 1 0 0 1 0-2h14a1 1 0 0 1 0 2Z" class="clr-i-solid--badged clr-i-solid-path-1--badged"/><circle cx="30" cy="5.75" r="5" fill="currentColor" class="clr-i-solid--badged clr-i-solid-path-2--badged clr-i-badge"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}