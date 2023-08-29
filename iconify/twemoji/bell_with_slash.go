package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BellWithSlash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#FFAC33" d="M28 13c0 11 5 10 5 15c0 0 0 2-2 2H5c-2 0-2-2-2-2c0-5 5-4 5-15C8 7.478 12.477 3 18 3s10 4.478 10 10z"/><circle cx="18" cy="3" r="3" fill="#FFAC33"/><path fill="#FFAC33" d="M18 36a4 4 0 0 0 4-4h-8a4 4 0 0 0 4 4z"/><path fill="#DD2E44" d="M31.003 33.003L3.997 5.997a1.41 1.41 0 0 1 0-1.994l.006-.006a1.41 1.41 0 0 1 1.994 0l27.006 27.006a1.41 1.41 0 0 1 0 1.994l-.006.006a1.41 1.41 0 0 1-1.994 0z"/>`),
		g.Group(children),
	)
}