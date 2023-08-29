package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Terminal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M472 40H40a24.028 24.028 0 0 0-24 24v392a24.028 24.028 0 0 0 24 24h432a24.028 24.028 0 0 0 24-24V64a24.028 24.028 0 0 0-24-24Zm-8 408H48V72h416Z"/><path fill="currentColor" d="m115.962 282.627l73.445-82.672l-73.451-82.588l-23.912 21.266l54.549 61.333l-54.555 61.407l23.924 21.254zM216 240h128v32H216z"/>`),
		g.Group(children),
	)
}