package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JaMobil(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m10.314 17.708l1.044 5.908a1.99 1.99 0 0 1-1.602 2.314h0a1.965 1.965 0 0 1-1.482-.332"/><circle cx="9.896" cy="15.344" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.856 19.902a2 2 0 0 1-1.62 2.318h0a2 2 0 0 1-2.319-1.621l-.226-1.28A2 2 0 0 1 14.312 17h0a2 2 0 0 1 2.318 1.621m.575 3.251l-.923-5.22"/><circle cx="19.418" cy="21.251" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m18.025 13.373l1.036 5.859"/><rect width="4" height="5.3" x="18.066" y="29.327" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2" ry="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.274 31.327a2 2 0 0 1 2-2h0a2 2 0 0 1 2 2v3.3m-4-5.3v5.3m4-3.3a2 2 0 0 1 2-2h0a2 2 0 0 1 2 2v3.3"/><circle cx="29.753" cy="26.877" r=".7" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.753 29.327v5.3m1.851-8v7a1 1 0 0 0 1 1h.3m-9.006-3.3a2 2 0 0 1 2-2h0a2 2 0 0 1 2 2v1.3a2 2 0 0 1-2 2h0a2 2 0 0 1-2-2m0 2v-8m10.613 0a2.907 2.907 0 0 0-2.908-2.908m5.431 2.908c0-3-2.431-5.431-5.43-5.431"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.726 26.627a8.122 8.122 0 0 0-8.122-8.123"/>`),
		g.Group(children),
	)
}