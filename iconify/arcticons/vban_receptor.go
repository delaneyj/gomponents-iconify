package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VbanReceptor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.966"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 18.101h37M9.251 38.95h1.967m-1.967-1.967h1.967m-1.967-1.967h1.967m1.966 0h1.967m-1.967 1.967h1.967m-1.967 1.967h1.967m-1.967-5.9h1.967m-1.967-1.967h1.967m1.966 0h1.967m-1.967 1.967h1.967m-1.967 1.966h1.967m-1.967 1.967h1.967m-1.967 1.967h1.967m1.966 0h1.967m-1.967-1.967h1.967m-1.967-1.967h1.967M21.05 33.05h1.967m-1.967-1.967h1.967m-1.967-1.966h1.967m1.966 0h1.967m-1.967-1.967h1.967m-1.967-1.966h1.967m1.966 0h1.967m-1.967-1.967h1.967m1.966 1.967h1.967m1.966 5.899h1.967m-1.967 1.967h1.967m-1.967 1.966h1.967m-1.967 1.967h1.967m-1.967 1.967h1.967m-5.9 0h1.967m-1.967-1.967h1.967m-1.967-1.967h1.967m-1.967-1.966h1.967m-1.967-1.967h1.967m-1.967-1.966h1.967m-1.967-1.967h1.967m-5.9 0h1.967m-1.967 1.967h1.967m-1.967 1.966h1.967m-1.967 1.967h1.967m-1.967 1.966h1.967m-1.967 1.967h1.967m-1.967 1.967h1.967m-5.9 0h1.967m-1.967-1.967h1.967m-1.967-1.967h1.967m-1.967-1.966h1.967m-1.967-1.967h1.967"/>`),
		g.Group(children),
	)
}