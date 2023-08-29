package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BackupRestoreLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M6 22h24v2H6z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M26 26h4v2h-4z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="M13 9.92L17 6v13a1 1 0 1 0 2 0V6l4 3.95a1 1 0 1 0 1.38-1.45L18 2.16L11.61 8.5A1 1 0 0 0 13 9.92Z" class="clr-i-outline clr-i-outline-path-3"/><path fill="currentColor" d="M30.84 13.37A1.94 1.94 0 0 0 28.93 12H21v2h7.95c1.05 2.94 2.77 7.65 3.05 8.48V30H4v-7.52C4.28 21.65 7.05 14 7.05 14H15v-2H7.07a1.92 1.92 0 0 0-1.9 1.32C2 22 2 22.1 2 22.33V30a2 2 0 0 0 2 2h28a2 2 0 0 0 2-2v-7.67c0-.23 0-.33-3.16-8.96Z" class="clr-i-outline clr-i-outline-path-4"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}