package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngryNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g fill="currentColor" fill-rule="evenodd" clip-path="url(#healthiconsAngryNegative0)" clip-rule="evenodd"><path d="M24 40c8.837 0 16-7.163 16-16S32.837 8 24 8S8 15.163 8 24s7.163 16 16 16Zm-.013-9c2.33 0 4.362.796 5.44 1.977l.075.079c.073.077.146.154.206.242c.379.565.593 2.003-.684 1.638c-4.654-1.214-5.116-1.26-9.862 0c-1.63.409-1.203-1.25-.756-1.797c1.037-1.27 3.147-2.139 5.58-2.139ZM17.5 28c1.38 0 2.5-1.79 2.5-4c0-.856-.168-1.65-.454-2.3c.398.233.753.479 1.05.735l1.307-1.513c-1.255-1.085-3.105-1.893-4.837-2.465a29.533 29.533 0 0 0-4.67-1.125l-.29 1.979a27.59 27.59 0 0 1 4.333 1.046l.019.006C15.598 20.995 15 22.386 15 24c0 2.21 1.12 4 2.5 4Zm13 0c1.38 0 2.5-1.79 2.5-4c0-1.714-.674-3.177-1.621-3.746a27.47 27.47 0 0 1 4.016-.943l-.29-1.979a29.5 29.5 0 0 0-4.67 1.125c-1.733.572-3.583 1.38-4.838 2.465l1.307 1.513c.458-.395 1.053-.765 1.733-1.103C28.241 22.04 28 22.975 28 24c0 2.21 1.12 4 2.5 4Z"/><path d="M0 0h48v48H0V0Zm24 42c9.941 0 18-8.059 18-18S33.941 6 24 6S6 14.059 6 24s8.059 18 18 18Z"/></g><defs><clipPath id="healthiconsAngryNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}