package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpLeftOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><g fill-rule="evenodd" clip-rule="evenodd"><g opacity=".2"><path d="M7.857 13.64A1.5 1.5 0 0 1 6.5 12.012l.472-5.185a1.5 1.5 0 0 1 2.987.271l-.471 5.186a1.5 1.5 0 0 1-1.63 1.358Z"/><path d="M15.144 6.354a1.5 1.5 0 0 1-1.358 1.63L8.6 8.455a1.5 1.5 0 0 1-.271-2.987l5.185-.472a1.5 1.5 0 0 1 1.63 1.358Z"/><path d="M8.818 7.315a1.5 1.5 0 0 1 2.121 0l5.657 5.657a1.5 1.5 0 0 1-2.121 2.121L8.818 9.437a1.5 1.5 0 0 1 0-2.122Z"/></g><path d="M5.948 12.145a.5.5 0 0 1-.453-.543l.471-5.186a.5.5 0 0 1 .996.09l-.471 5.186a.5.5 0 0 1-.543.453Z"/><path d="M12.148 5.945a.5.5 0 0 1-.453.543L6.51 6.96a.5.5 0 0 1-.09-.996l5.185-.472a.5.5 0 0 1 .543.453Z"/><path d="M6.647 6.643a.5.5 0 0 1 .707 0l6.535 6.536a.5.5 0 1 1-.707.707L6.646 7.351a.5.5 0 0 1 0-.708Z"/></g><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}