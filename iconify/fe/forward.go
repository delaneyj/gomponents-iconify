package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Forward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feForward0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feForward1" fill="currentColor" fill-rule="nonzero"><path id="feForward2" d="M12 12V7.65c0-.11.025-.22.072-.316c.152-.314.5-.428.775-.253l6.86 4.349c.093.059.17.147.221.253c.153.314.054.71-.221.885l-6.86 4.35a.516.516 0 0 1-.277.081c-.315 0-.57-.291-.57-.651v-4.35c0 .23-.106.451-.293.57l-6.86 4.35A.516.516 0 0 1 4.57 17c-.315 0-.57-.291-.57-.651V7.651c0-.11.025-.22.072-.316c.152-.314.5-.428.775-.253l6.86 4.349c.093.059.17.147.221.253c.049.1.072.209.072.315Z"/></g></g>`),
		g.Group(children),
	)
}