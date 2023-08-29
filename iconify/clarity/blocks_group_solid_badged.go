package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlocksGroupSolidBadged(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="m33.53 18.76l-6.93-3.19v-2.88a7.66 7.66 0 0 1-2-1.47v4.34l-6.43 3v-7.69a1.05 1.05 0 0 0 .35-.08L23 8.73a7.65 7.65 0 0 1-.48-2l-4.42 2l-5.1-2.3l5.1-2.35l4.38 2V6a7.55 7.55 0 0 1 .27-2L18.5 2.08a1 1 0 0 0-.84 0l-7.5 3.45a1 1 0 0 0-.58.91v9.14l-6.9 3.18a1 1 0 0 0-.58.91v9.78a1 1 0 0 0 .58.91l7.5 3.45a1 1 0 0 0 .84 0l7.08-3.26l7.08 3.26a1 1 0 0 0 .84 0l7.5-3.45a1 1 0 0 0 .58-.91v-9.78a1 1 0 0 0-.57-.91Zm-28 .91l5.11-2.36l5.11 2.35L10.6 22ZM10.6 31.8v-7.69A1.08 1.08 0 0 0 11 24l6.08-2.8v7.6Zm9.9-12.13l5.11-2.35l5.11 2.35L25.61 22Zm5.14 12.13v-7.69A.89.89 0 0 0 26 24l6.08-2.8v7.6Z" class="clr-i-solid--badged clr-i-solid-path-1--badged"/><circle cx="29.98" cy="6" r="5" fill="currentColor" class="clr-i-solid--badged clr-i-solid-path-2--badged clr-i-badge"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}