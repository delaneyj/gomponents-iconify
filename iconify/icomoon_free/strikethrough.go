package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Strikethrough(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M16 8v1h-3.664c.43.602.664 1.292.664 2c0 1.107-.573 2.172-1.572 2.921C10.501 14.617 9.283 15 8 15s-2.501-.383-3.428-1.079C3.573 13.172 3 12.107 3 11h2c0 1.084 1.374 2 3 2s3-.916 3-2s-1.374-2-3-2H0V8h4.68a3.003 3.003 0 0 1-.108-.079C3.573 7.172 3 6.107 3 5s.573-2.172 1.572-2.921C5.499 1.383 6.717 1 8 1s2.501.383 3.428 1.079C12.427 2.828 13 3.893 13 5h-2c0-1.084-1.374-2-3-2s-3 .916-3 2s1.374 2 3 2c1.234 0 2.407.354 3.32 1H16z"/>`),
		g.Group(children),
	)
}