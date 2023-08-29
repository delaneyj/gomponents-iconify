package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unifiednlp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.9 3.5a1 1 0 0 0-1 .86l-.77 5.43a15.36 15.36 0 0 0-3.46 2L9.54 9.75a1 1 0 0 0-1.25.44l-4.1 7.11a1 1 0 0 0 .25 1.31L8.76 22a19 19 0 0 0-.14 2a16.52 16.52 0 0 0 .14 2l-4.32 3.39a1 1 0 0 0-.25 1.31l4.1 7.11a1 1 0 0 0 1.25.44l5.11-2.06a15.68 15.68 0 0 0 3.46 2l.77 5.43a1 1 0 0 0 1 .86h8.2a1 1 0 0 0 1-.86l.77-5.43a15.36 15.36 0 0 0 3.46-2l5.11 2.06a1 1 0 0 0 1.25-.44l4.1-7.11a1 1 0 0 0-.25-1.31L39.23 26a16.52 16.52 0 0 0 .14-2a16.52 16.52 0 0 0-.14-2l4.33-3.39a1 1 0 0 0 .25-1.31l-4.1-7.11a1 1 0 0 0-1.25-.44l-5.11 2.06a15.68 15.68 0 0 0-3.46-2l-.77-5.43a1 1 0 0 0-1-.86Zm4.1 9.92a8 8 0 0 1 8 8c0 4.68-3.8 8.33-8 13.12c-4.06-4.58-7.72-8.13-8-12.53v0a5.06 5.06 0 0 1 0-.55a8 8 0 0 1 8-8Zm0 4.18a3.87 3.87 0 1 0 3.87 3.86A3.86 3.86 0 0 0 24 17.6Z"/>`),
		g.Group(children),
	)
}