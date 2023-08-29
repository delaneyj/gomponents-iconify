package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoCameraSlashFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M213.92 210.62a8 8 0 1 1-11.84 10.76L182.64 200H32a16 16 0 0 1-16-16V72a16 16 0 0 1 16-16h19.73l-9.65-10.62a8 8 0 1 1 11.84-10.76ZM251.77 73a8 8 0 0 0-8.21.39l-32 21.34a8 8 0 0 0-3.56 6.65v53.34a8 8 0 0 0 3.56 6.65l32 21.34A8 8 0 0 0 248 184a8 8 0 0 0 8-8V80a8 8 0 0 0-4.23-7Zm-73.69 74.46A8 8 0 0 0 192 142V72a16 16 0 0 0-16-16h-62.94a8 8 0 0 0-5.92 13.38Z"/>`),
		g.Group(children),
	)
}