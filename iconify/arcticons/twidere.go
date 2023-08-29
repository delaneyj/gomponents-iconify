package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Twidere(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m23.57 14.64l-6.32 6.42l-6.33-6.42l6.33-6.37l6.32 6.37z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 8.27h12.76l-6.33 6.37L4.5 8.27zm6.42 6.37l6.33 6.42l-6.35 6.34l.02-12.76z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m17.2 39.72l-9.24-9.36l9.29-9.3l-.05 18.66zm.05.01V21.06l18.59 18.6l-18.59.07z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.4 34.22L17.25 21.06H43.5L30.4 34.22z"/>`),
		g.Group(children),
	)
}