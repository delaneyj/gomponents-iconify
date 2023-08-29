package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spotube(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.17 24.589a4.31 4.31 0 0 0-4.659 4.607v5.858s-.466 3.997 4.659 3.997V23.485a14.83 14.83 0 0 1 29.66 0V39.05c5.122 0 4.66-3.997 4.66-3.997v-5.858a4.303 4.303 0 0 0-4.66-4.607m-26.967.632v13.778"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m14.387 27.955l5.543.885l-1.988 6.29l-3.555.059m19.142-7.234l-5.133.885l1.81 6.08l3.323.205m-15.356-.65s3.703 4.439 5.522 4.87m-1.209-3.818a11.71 11.71 0 0 0 3.04 3.124m-1.463-4.281a16.297 16.297 0 0 0 3.23 3.418m1.397-1.104a24.03 24.03 0 0 0-4.564-4.607m3.829-2.545s-3.597-.649-4.607-.03s-3.092 2.304-2.335 3.103m6.942 3.133l1.746-1.42m6.352-9.098v13.778"/>`),
		g.Group(children),
	)
}