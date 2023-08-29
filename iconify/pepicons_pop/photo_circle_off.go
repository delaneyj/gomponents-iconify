package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhotoCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M11 12a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z" clip-rule="evenodd"/><path d="M9.697 13.082a1 1 0 0 1 1.606 0l1.91 2.572a1 1 0 0 1-.804 1.596H8.591a1 1 0 0 1-.803-1.596l1.909-2.572Z"/><path fill-rule="evenodd" d="m10.5 14.517l-.916 1.233h1.832l-.916-1.233Zm.803-1.435a1 1 0 0 0-1.606 0l-1.91 2.572a1 1 0 0 0 .804 1.596h3.818a1 1 0 0 0 .803-1.596l-1.91-2.572Z" clip-rule="evenodd"/><path d="M13.494 12.123a1 1 0 0 1 1.512 0l3.007 3.472a1 1 0 0 1-.756 1.655h-6.014a1 1 0 0 1-.756-1.655l3.007-3.472Z"/><path fill-rule="evenodd" d="m14.25 13.541l-1.913 2.209h3.826l-1.913-2.209Zm.756-1.418a1 1 0 0 0-1.512 0l-3.007 3.472a1 1 0 0 0 .756 1.655h6.014a1 1 0 0 0 .756-1.655l-3.007-3.472Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M6.5 6.5v13h13v-13h-13Zm-.5-2A1.5 1.5 0 0 0 4.5 6v14A1.5 1.5 0 0 0 6 21.5h14a1.5 1.5 0 0 0 1.5-1.5V6A1.5 1.5 0 0 0 20 4.5H6Z" clip-rule="evenodd"/><path d="M4.293 5.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/><path fill-rule="evenodd" d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}