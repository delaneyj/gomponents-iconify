package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Borrowed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.694 41.287H20.563c-1.82 0-3.336-1.516-3.336-3.336V24h-8.39C7.016 24 5.5 22.484 5.5 20.664s1.516-3.336 3.336-3.336h8.39v-7.279c0-1.82 1.517-3.336 3.337-3.336s3.336 1.516 3.336 3.336v7.279h8.795c5.661 0 9.806 5.054 9.806 12.03c0 6.874-4.044 11.929-9.806 11.929Zm-8.795-6.672h8.795c.607 0 1.314-.101 2.123-1.213c.708-1.011 1.01-2.528 1.01-4.145c0-2.123-.808-5.358-3.133-5.358h-8.795v10.716Z"/>`),
		g.Group(children),
	)
}