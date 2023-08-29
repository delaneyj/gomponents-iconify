package ei

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScGithub(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M25 10c-8.3 0-15 6.7-15 15c0 6.6 4.3 12.2 10.3 14.2c.8.1 1-.3 1-.7v-2.6c-4.2.9-5.1-2-5.1-2c-.7-1.7-1.7-2.2-1.7-2.2c-1.4-.9.1-.9.1-.9c1.5.1 2.3 1.5 2.3 1.5c1.3 2.3 3.5 1.6 4.4 1.2c.1-1 .5-1.6 1-2c-3.3-.4-6.8-1.7-6.8-7.4c0-1.6.6-3 1.5-4c-.2-.4-.7-1.9.1-4c0 0 1.3-.4 4.1 1.5c1.2-.3 2.5-.5 3.8-.5c1.3 0 2.6.2 3.8.5c2.9-1.9 4.1-1.5 4.1-1.5c.8 2.1.3 3.6.1 4c1 1 1.5 2.4 1.5 4c0 5.8-3.5 7-6.8 7.4c.5.5 1 1.4 1 2.8v4.1c0 .4.3.9 1 .7c6-2 10.2-7.6 10.2-14.2C40 16.7 33.3 10 25 10z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}