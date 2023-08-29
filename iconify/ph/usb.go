package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Usb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m252.44 121.34l-48-32A8 8 0 0 0 192 96v24H72V72h33a32 32 0 1 0 0-16H72a16 16 0 0 0-16 16v48H8a8 8 0 0 0 0 16h48v48a16 16 0 0 0 16 16h32v8a16 16 0 0 0 16 16h32a16 16 0 0 0 16-16v-32a16 16 0 0 0-16-16h-32a16 16 0 0 0-16 16v8H72v-48h120v24a8 8 0 0 0 12.44 6.66l48-32a8 8 0 0 0 0-13.32ZM136 48a16 16 0 1 1-16 16a16 16 0 0 1 16-16Zm-16 128h32v32h-32Zm88-30.95V111l25.58 17Z"/>`),
		g.Group(children),
	)
}