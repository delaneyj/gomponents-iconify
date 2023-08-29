package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Girl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M10 12a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm5 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/><path fill-rule="evenodd" d="M12.024 2H12C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10c0-5.258-4.058-9.568-9.212-9.97v-.002A9.94 9.94 0 0 0 12.025 2ZM12 20a8 8 0 0 0 7.742-10.022a10.016 10.016 0 0 1-8.982-4.376a7.976 7.976 0 0 1-5.691 2.4A8 8 0 0 0 12 20Zm-.021-16h.045h-.045Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}