package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DigitZero(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#4fd1d9"/><path fill="#fff" d="M32 16c3 0 5.3 1.1 7 3.2c2 2.6 3 6.8 3 12.8c0 5.9-1 10.2-3 12.8c-1.7 2.1-4 3.2-7 3.2s-5.4-1.2-7.2-3.5C22.9 42.1 22 38 22 31.9c0-5.9 1-10.1 3-12.7c1.7-2.1 4-3.2 7-3.2m0 5c-.7 0-1.4.2-1.9.7s-1 1.3-1.3 2.5c-.4 1.6-.6 4.2-.6 7.8c0 3.7.2 6.2.5 7.6c.4 1.4.8 2.3 1.4 2.7c.6.5 1.2.7 1.9.7s1.4-.2 1.9-.7c.6-.5 1-1.3 1.3-2.5c.4-1.5.6-4.1.6-7.8s-.2-6.2-.5-7.6c-.4-1.4-.8-2.3-1.4-2.8c-.6-.4-1.2-.6-1.9-.6"/>`),
		g.Group(children),
	)
}