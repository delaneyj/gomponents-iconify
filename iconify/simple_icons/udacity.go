package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Udacity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 0L0 4.8v9.6C0 20.8 4.8 24 8.8 24c1.348 0 2.786-.362 4.1-1.088l6.303-3.633C21.687 18.155 24 15.64 24 11.2V.8L22.4 0L16 4v10.4c0 1.6-.3 2.898-.785 3.948c-2.002-.257-5.615-1.597-5.615-7.15V.802zm0 1.76v9.44c0 5.342 3.346 7.9 6.313 8.597c-1.618 1.99-4.025 2.603-5.512 2.603c-2.4 0-7.2-1.6-7.2-8V5.6zm14.4.04v9.4c0 5.45-3.482 6.84-5.504 7.132c.446-1.14.704-2.45.704-3.932V4.8z"/>`),
		g.Group(children),
	)
}