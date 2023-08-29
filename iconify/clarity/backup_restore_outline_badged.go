package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BackupRestoreOutlineBadged(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M6 22h24v2H6z" class="clr-i-outline--badged clr-i-outline-path-1--badged"/><path fill="currentColor" d="M26 26h4v2h-4z" class="clr-i-outline--badged clr-i-outline-path-2--badged"/><path fill="currentColor" d="M13 9.92L17 6v13a1 1 0 1 0 2 0V6l4 3.95a1 1 0 0 0 .71.29h.11a7.46 7.46 0 0 1-1.25-3.52L18 2.16L11.61 8.5A1 1 0 0 0 13 9.92Z" class="clr-i-outline--badged clr-i-outline-path-3--badged"/><path fill="currentColor" d="M30.87 13.45a7.55 7.55 0 0 1-.87.05a7.46 7.46 0 0 1-4.49-1.5H21v2h7.95c1.05 2.94 2.77 7.65 3.05 8.48V30H4v-7.52C4.28 21.65 7.05 14 7.05 14H15v-2H7.07a1.92 1.92 0 0 0-1.9 1.32C2 22 2 22.1 2 22.33V30a2 2 0 0 0 2 2h28a2 2 0 0 0 2-2v-7.67c0-.23 0-.33-3.13-8.88Z" class="clr-i-outline--badged clr-i-outline-path-4--badged"/><circle cx="30" cy="6" r="5" fill="currentColor" class="clr-i-outline--badged clr-i-outline-path-5--badged clr-i-badge"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}