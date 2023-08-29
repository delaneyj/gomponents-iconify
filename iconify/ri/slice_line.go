package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SliceLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15.69 12.915l1.769 1.767c-6.01 6.01-10.96 6.01-15.203 4.597L17.812 3.722l3.536 3.536l-5.657 5.657Zm-2.827 0l5.656-5.657l-.707-.707L6.314 18.048c2.732.108 5.358-.906 8.267-3.415l-1.718-1.718Z"/>`),
		g.Group(children),
	)
}