package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Camera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M471.993 112h-89.2l-16.242-46.75a32.023 32.023 0 0 0-30.229-21.5H175.241a31.991 31.991 0 0 0-30.294 21.691L129.1 112H40a24.027 24.027 0 0 0-24 24v312a24.027 24.027 0 0 0 24 24h431.993a24.027 24.027 0 0 0 24-24V136a24.027 24.027 0 0 0-24-24Zm-8 328H48.007V144h104.01l23.224-68.25h161.081l23.71 68.25h103.961Z"/><path fill="currentColor" d="M256 168a114 114 0 1 0 114 114a114.13 114.13 0 0 0-114-114Zm0 196a82 82 0 1 1 82-82a82.093 82.093 0 0 1-82 82Z"/>`),
		g.Group(children),
	)
}