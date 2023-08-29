package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlocksGroupSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="m33.53 18.76l-6.93-3.19V6.43a1 1 0 0 0-.6-.9l-7.5-3.45a1 1 0 0 0-.84 0l-7.5 3.45a1 1 0 0 0-.58.91v9.14l-6.9 3.18a1 1 0 0 0-.58.91v9.78a1 1 0 0 0 .58.91l7.5 3.45a1 1 0 0 0 .84 0l7.08-3.26l7.08 3.26a1 1 0 0 0 .84 0l7.5-3.45a1 1 0 0 0 .58-.91v-9.78a1 1 0 0 0-.57-.91ZM25.61 22l-5.11-2.33l5.11-2.35l5.11 2.35Zm-1-6.44l-6.44 3v-7.69a1 1 0 0 0 .35-.08L24.6 8v7.58ZM18.1 4.08l5.11 2.35l-5.11 2.35L13 6.43Zm-7.5 13.23l5.11 2.35L10.6 22l-5.11-2.33Zm6.5 11.49l-6.5 3v-7.69A1 1 0 0 0 11 24l6.08-2.8Zm15 0l-6.46 3v-7.69A1 1 0 0 0 26 24l6.08-2.8Z" class="clr-i-solid clr-i-solid-path-1"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}