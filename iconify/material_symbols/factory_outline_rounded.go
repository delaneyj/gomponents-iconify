package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FactoryOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 22q-.825 0-1.413-.588T2 20v-8.7q0-.6.325-1.1t.9-.75L7.6 7.6q.5-.2.95.075T9 8.5V9l3.625-1.45q.5-.2.937.1t.438.825V10h8v10q0 .825-.588 1.413T20 22H4Zm0-2h16v-8h-8V9.95l-5 2V10l-3 1.325V20Zm7-2h2v-4h-2v4Zm-4 0h2v-4H7v4Zm8 0h2v-4h-2v4Zm7-8h-5l.9-7.125q.05-.375.325-.625t.65-.25h1.25q.375 0 .65.25t.325.625L22 10ZM4 20h16H4Z"/>`),
		g.Group(children),
	)
}