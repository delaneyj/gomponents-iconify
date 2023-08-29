package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trifa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.939 18.753c2.12-2.646 2.727-6.17.783-8.946a6.985 6.985 0 1 0-5.113 10.965a10.601 10.601 0 0 1-.597 2.699a28.957 28.957 0 0 0 4.927-4.718Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.687 13.813a9.313 9.313 0 1 1 18.626 0v5.821h6.403V43.5H8.284V19.634h6.403Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.015 33.022a6.985 6.985 0 1 1 13.97 0v6.636h-13.97Z"/><circle cx="20.799" cy="30.403" r=".75" fill="currentColor"/><circle cx="27.201" cy="30.403" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m19.925 27.201l-1.164-1.746m9.314 1.746l1.164-1.746"/>`),
		g.Group(children),
	)
}