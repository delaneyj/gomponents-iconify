package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Read(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 18.883a10.8 10.8 0 0 1-9.675-5.728a2.6 2.6 0 0 1 0-2.31A10.8 10.8 0 0 1 12 5.117a10.8 10.8 0 0 1 9.675 5.728a2.6 2.6 0 0 1 0 2.31A10.8 10.8 0 0 1 12 18.883Zm0-12.766a9.787 9.787 0 0 0-8.78 5.176a1.586 1.586 0 0 0 0 1.415A9.788 9.788 0 0 0 12 17.883a9.787 9.787 0 0 0 8.78-5.176a1.584 1.584 0 0 0 0-1.414A9.787 9.787 0 0 0 12 6.117Z"/><path fill="currentColor" d="M12 16.049A4.049 4.049 0 1 1 16.049 12A4.054 4.054 0 0 1 12 16.049Zm0-7.1A3.049 3.049 0 1 0 15.049 12A3.052 3.052 0 0 0 12 8.951Z"/><circle cx="12" cy="12" r="2.028" fill="currentColor"/>`),
		g.Group(children),
	)
}