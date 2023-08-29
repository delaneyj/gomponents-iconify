package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Superfreezz(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.55 5.5h-5.66L31 10.13h7.76L24.38 25.22h7l-19 13.68l8-11.46h-4.87l11.39-15h-7.25L25 5.5H7.45a2 2 0 0 0-1.95 2v33.1a2 2 0 0 0 2 2h33.1a2 2 0 0 0 2-2V7.45a2 2 0 0 0-2.05-1.95Zm-6.74 19.81h5.34m-7.02 5.35h3.49m-1.81-5.35l-3.35 10.71"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.69 12.47a4 4 0 0 0-3.51 2.68h0A1.9 1.9 0 0 0 12 17.83h.91m.02 0h.9a1.9 1.9 0 0 1 1.84 2.67h0a4 4 0 0 1-3.51 2.68m5.48-9.81c-.54-.61-1.25-.9-3-.9h-.91m-5.52 9.81c.54.62 1.25.9 3 .9h.91"/>`),
		g.Group(children),
	)
}