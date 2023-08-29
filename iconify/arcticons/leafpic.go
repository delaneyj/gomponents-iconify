package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Leafpic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m18.5 9.1l1.5 2.09l2.54-.42l-.43 2.54l2.09 1.49l-2.1 1.49l.43 2.54l-2.53-.42l-1.5 2.09l-1.5-2.09l-2.53.42l.42-2.53l-2.09-1.5l2.09-1.49l-.42-2.53l2.53.41ZM37.09 25v11.34A10.94 10.94 0 0 1 32.77 43a7.85 7.85 0 0 1-.76.52a12.72 12.72 0 0 1-1.31.71a13.33 13.33 0 0 1-1.52.6a17.12 17.12 0 0 1-1.74.46c-.65.13-1.31.23-2 .3s-1.48.1-2.22.1q-.78 0-1.5-.06h-.47a14.78 14.78 0 0 1-1.74-.32a13.67 13.67 0 0 1-1.53-.5l-.69-.31l-.64-.34a11 11 0 0 1-1.15-.77l-.4-.33a12.09 12.09 0 0 1-3.82-6.85v-.08c0-.17-.06-.35-.08-.53V28c8.51 0 11 5.7 11.76 8.35h0A13.65 13.65 0 0 1 37.09 25Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.77 43H41a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2v34a2 2 0 0 0 2 2h8.13"/>`),
		g.Group(children),
	)
}