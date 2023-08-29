package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Newsprovider(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.29 42.5A28.79 28.79 0 0 0 5.5 13.71V42.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.395 38.395a28.79 28.79 0 0 0-28.79-28.79v28.79Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.5 34.29A28.79 28.79 0 0 0 13.71 5.5v28.79Z"/><circle cx="19.322" cy="28.678" r="2.876" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}