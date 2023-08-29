package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PushPinSimpleSlashFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M213.38 221.92a8 8 0 0 1-11.3-.54L168.1 184H136v56a8 8 0 0 1-16 0v-56H40a8 8 0 0 1 0-16h9.29l16.95-96l-24.16-26.62a8 8 0 1 1 11.84-10.76l160 176a8 8 0 0 1-.54 11.3Zm-22.88-60.85a8 8 0 0 0 13.79-6.77L185.54 48H192a8 8 0 0 0 0-16H91.25a8 8 0 0 0-5.92 13.38Z"/>`),
		g.Group(children),
	)
}