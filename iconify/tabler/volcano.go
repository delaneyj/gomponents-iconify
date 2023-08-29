package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Volcano(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 8V7a2 2 0 1 0-4 0m10 1V7a2 2 0 1 1 4 0M4 20l3.472-7.812A2 2 0 0 1 9.3 11h5.4a2 2 0 0 1 1.828 1.188L20 20"/><path d="M6.192 15.064A2.14 2.14 0 0 1 6.667 15c.527-.009 1.026.178 1.333.5c.307.32.806.507 1.333.5c.527.007 1.026-.18 1.334-.5c.307-.322.806-.509 1.333-.5c.527-.009 1.026.178 1.333.5c.308.32.807.507 1.334.5c.527.007 1.026-.18 1.333-.5c.307-.322.806-.509 1.333-.5c.161.003.32.025.472.064M12 8V4"/></g>`),
		g.Group(children),
	)
}