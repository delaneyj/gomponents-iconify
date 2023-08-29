package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Keypad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 400a48 48 0 1 0 48 48a48 48 0 0 0-48-48Zm0-128a48 48 0 1 0 48 48a48 48 0 0 0-48-48Zm0-128a48 48 0 1 0 48 48a48 48 0 0 0-48-48Zm0-128a48 48 0 1 0 48 48a48 48 0 0 0-48-48Zm128 256a48 48 0 1 0 48 48a48 48 0 0 0-48-48Zm0-128a48 48 0 1 0 48 48a48 48 0 0 0-48-48Zm0-128a48 48 0 1 0 48 48a48 48 0 0 0-48-48ZM128 272a48 48 0 1 0 48 48a48 48 0 0 0-48-48Zm0-128a48 48 0 1 0 48 48a48 48 0 0 0-48-48Zm0-128a48 48 0 1 0 48 48a48 48 0 0 0-48-48Z"/>`),
		g.Group(children),
	)
}