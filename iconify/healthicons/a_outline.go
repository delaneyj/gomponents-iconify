package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M21.23 10.846a3 3 0 0 1 5.54 0l7.478 17.95c.015.033.03.068.043.102l2.478 5.948a3 3 0 1 1-5.538 2.308L29.5 33h-11l-1.73 4.154a3 3 0 1 1-5.54-2.308l2.48-5.946l.043-.104l7.479-17.949ZM24 11a1 1 0 0 0-.923.615l-7.483 17.96l-.009.02a1.021 1.021 0 0 0-.017.042l-.009.021l-2.482 5.957a1 1 0 1 0 1.846.77l1.987-4.77a1 1 0 0 1 .923-.615h12.334a1 1 0 0 1 .923.615l1.987 4.77a1 1 0 0 0 1.846-.77l-2.482-5.957l-.009-.021l-.009-.023a1.488 1.488 0 0 0-.008-.019l-.01-.02l-7.482-17.96A1 1 0 0 0 24 11Zm0 5.2a1 1 0 0 1 .923.615l4.5 10.8A1 1 0 0 1 28.5 29h-9a1 1 0 0 1-.923-1.385l4.5-10.8A1 1 0 0 1 24 16.2ZM21 27h6l-3-7.2l-3 7.2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}