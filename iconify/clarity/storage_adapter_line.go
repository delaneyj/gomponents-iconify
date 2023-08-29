package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StorageAdapterLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M6.06 30a1 1 0 0 1-1-1V8h-2a1 1 0 0 1 0-2h4v23a1 1 0 0 1-1 1Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M30.06 27h-25V9h25a3 3 0 0 1 3 3v12a3 3 0 0 1-3 3Zm-23-2h23a1 1 0 0 0 1-1V12a1 1 0 0 0-1-1h-23Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="M22.06 20h6v2h-6z" class="clr-i-outline clr-i-outline-path-3"/><path fill="currentColor" d="M22.06 14h6v2h-6z" class="clr-i-outline clr-i-outline-path-4"/><path fill="currentColor" d="M19.06 22h-8v-2h7v-6h2v7a1 1 0 0 1-1 1Z" class="clr-i-outline clr-i-outline-path-5"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}