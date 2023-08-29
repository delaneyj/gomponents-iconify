package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Exportcontacts(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M19.9 3.5a1 1 0 0 0-1 .86l-.77 5.43a15 15 0 0 0-3.46 2L9.54 9.75a1 1 0 0 0-1.25.45l-4.1 7.1a1 1 0 0 0 .25 1.31L8.76 22a16.66 16.66 0 0 0-.14 2a16.52 16.52 0 0 0 .14 2l-4.32 3.39a1 1 0 0 0-.25 1.31l4.1 7.11a1 1 0 0 0 1.25.44l5.11-2.06a15.26 15.26 0 0 0 3.46 2l.77 5.44a1 1 0 0 0 1 .86h8.19a1 1 0 0 0 1-.86l.77-5.43a15.36 15.36 0 0 0 3.46-2l5.11 2.06a1 1 0 0 0 1.24-.45l4.11-7.1a1 1 0 0 0-.25-1.31L39.22 26a14.83 14.83 0 0 0 .15-2a14.59 14.59 0 0 0-.15-2l4.34-3.39a1 1 0 0 0 .24-1.31l-4.1-7.11a1 1 0 0 0-1.25-.44l-5.1 2.06a15.58 15.58 0 0 0-3.46-2l-.77-5.43a1.07 1.07 0 0 0-1-.86H19.9Zm4.1 9.41a4.65 4.65 0 1 1 0 9.3a4.65 4.65 0 1 1 0-9.3Zm0 11c5.18 0 9.31 1.44 9.31 3.17v4.44H14.69v-4.44c0-1.73 4.13-3.18 9.31-3.18Z"/>`),
		g.Group(children),
	)
}