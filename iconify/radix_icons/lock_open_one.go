package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockOpenOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.499 0C6.326 0 5.36.39 4.738 1.194C4.238 1.839 4 2.682 4 3.634h1c0-.79.197-1.4.528-1.828c.388-.5 1.024-.806 1.97-.806c.859 0 1.465.265 1.86.686c.4.426.642 1.074.642 1.95V6H3a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h9a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1h-1V3.636c0-1.055-.293-1.974-.912-2.634C9.465.338 8.57 0 7.498 0ZM3 7h9v6H3V7Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}