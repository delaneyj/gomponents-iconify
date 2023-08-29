package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tinyscanner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.2 20.38c.18.1 4.88 3.56 9.16 6.67c5.24 3.82 7.81 5.77 7.88 6a17.71 17.71 0 0 1 0 2.65c-.12 2.69-.18 2.82-1.44 3.12c-.56.13-.74.24-.8.48c-.15.59-.38.67-4.09 1.42l-4 .84c-7.73 1.62-7.95 1.66-9.07 1.65a5.21 5.21 0 0 1-1.6-.21a6.64 6.64 0 0 1-1.94-1.47l-5.66-5.23c-4.17-3.83-4.88-4.55-5-5a2.36 2.36 0 0 0-.8-1.15l-.71-.68v-2.83c0-2.81 0-2.84.26-2.93m0 0a2.35 2.35 0 0 1-1.51-2.89c.09-.16 1-1.54 2.1-3.07l4.36-6.23C15 7.71 15.47 7.18 16.41 6.7c.71-.36 1.16-.44 7.24-1.22l7.11-.92a31 31 0 0 1 3.81-.35h0c.52.1 1.48.94 1.48 1.3s-.64 1.18-6.17 8.6c-2.18 2.89-4.52 6.13-4.68 6.27"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m42.24 33l-19.53 4.36a.94.94 0 0 1-.85-.23L7.19 23.88m.2-.13l17.81-3.37"/>`),
		g.Group(children),
	)
}