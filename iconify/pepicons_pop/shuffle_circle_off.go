package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShuffleCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M6 8.75a1 1 0 0 1 1-1c3.335 0 6.5 2.126 6.5 5.25s-3.165 5.25-6.5 5.25a1 1 0 1 1 0-2c2.74 0 4.5-1.68 4.5-3.25S9.74 9.75 7 9.75a1 1 0 0 1-1-1Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M19 8.75a1 1 0 0 0-1-1c-3.335 0-6.5 2.126-6.5 5.25s3.165 5.25 6.5 5.25a1 1 0 1 0 0-2c-2.74 0-4.5-1.68-4.5-3.25s1.76-3.25 4.5-3.25a1 1 0 0 0 1-1Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M19.46 17.778a.75.75 0 0 1-1.053-.133l-1.5-1.936a.75.75 0 0 1 1.186-.918l1.5 1.935a.75.75 0 0 1-.134 1.052Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M17.05 19.785a.75.75 0 0 1-.15-1.05l1.5-2a.75.75 0 1 1 1.2.9l-1.5 2a.75.75 0 0 1-1.05.15Zm2.41-11.563a.75.75 0 0 0-1.053.133l-1.5 1.936a.75.75 0 0 0 1.186.918l1.5-1.935a.75.75 0 0 0-.134-1.052Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M17.05 6.15a.75.75 0 0 0-.15 1.05l1.5 2a.75.75 0 1 0 1.2-.9l-1.5-2a.75.75 0 0 0-1.05-.15Z" clip-rule="evenodd"/><path d="M4.293 5.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/><path fill-rule="evenodd" d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}