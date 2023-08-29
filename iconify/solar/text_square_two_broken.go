package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextSquareTwoBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M12 3C7.749 3 5.623 3 4.303 4.318C3.298 5.321 3.058 6.788 3 9.3m18 0c-.058-2.512-.298-3.98-1.303-4.982c-.818-.817-1.946-1.127-3.697-1.246M21 14.7c-.058 2.512-.298 3.98-1.303 4.982C18.377 21 16.251 21 12 21c-4.251 0-6.377 0-7.697-1.318C3.298 18.679 3.058 17.212 3 14.7M8 8h8m-4 8V8m10 4h-2M4 12H2"/>`),
		g.Group(children),
	)
}