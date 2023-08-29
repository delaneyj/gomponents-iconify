package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Linkbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><rect width="24.295" height="15.857" x="11.835" y="16.087" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="7.928"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.894 22.206h4.173a1.77 1.77 0 0 1 1.77 1.77h0a1.77 1.77 0 0 1-1.77 1.768h-4.173a1.77 1.77 0 0 1-1.769-1.769h0a1.77 1.77 0 0 1 1.77-1.769Z"/>`),
		g.Group(children),
	)
}