package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GovBr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 42.5h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2Z"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m26.664 21.062l-2.218 5.876l-2.217-5.876m-8.294 0v6.652a2.217 2.217 0 0 1-2.218 2.217h0a2.21 2.21 0 0 1-1.568-.65"/><rect width="4.435" height="5.876" x="9.5" y="21.062" rx="2.217" ry="2.217" transform="rotate(-180 11.717 24)"/><rect width="4.435" height="5.876" x="16.085" y="21.062" rx="2.217" ry="2.217"/><path d="M36.283 23.28c0-1.225.992-2.218 2.217-2.218h0m-2.217 0v5.876m-6.58-3.658c0-1.225.993-2.218 2.218-2.218h0c1.224 0 2.217.993 2.217 2.217v1.442a2.217 2.217 0 0 1-2.217 2.217h0a2.217 2.217 0 0 1-2.218-2.217m0 2.217v-8.869"/></g><circle cx="27.265" cy="26.724" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}