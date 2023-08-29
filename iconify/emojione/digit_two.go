package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DigitTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#4fd1d9"/><path fill="#fff" d="M42 42.3V48H22c.2-2.1.9-4.2 1.9-6.1c1.1-1.9 3.2-4.5 6.4-7.6c2.6-2.6 4.1-4.3 4.7-5.2c.8-1.3 1.2-2.5 1.2-3.7c0-1.4-.3-2.4-1-3.1s-1.6-1.1-2.8-1.1c-1.2 0-2.1.4-2.8 1.2c-.7.8-1.1 2-1.2 3.8l-5.7-.6c.3-3.4 1.4-5.8 3.2-7.2c1.8-1.5 4-2.2 6.7-2.2c2.9 0 5.2.8 6.9 2.5S42 22.5 42 25c0 1.4-.2 2.8-.7 4.1c-.5 1.2-1.3 2.5-2.3 3.9c-.7.9-1.9 2.3-3.7 4.1c-1.8 1.8-2.9 2.9-3.4 3.5c-.5.6-.9 1.1-1.2 1.7H42"/>`),
		g.Group(children),
	)
}