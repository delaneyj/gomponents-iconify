package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceSmilingCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopFaceSmilingCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M13 20.5a7.5 7.5 0 1 0 0-15a7.5 7.5 0 0 0 0 15Zm0 2a9.5 9.5 0 1 0 0-19a9.5 9.5 0 0 0 0 19Z" clip-rule="evenodd"/><path d="M11.5 10.25a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0Zm5.5 0a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0Z"/><path fill-rule="evenodd" d="M17.447 15.105a1 1 0 0 1 .447 1.342L17 16l.894.448v.001l-.002.003l-.003.005l-.006.013a1.849 1.849 0 0 1-.077.135a4.04 4.04 0 0 1-.21.308a4.842 4.842 0 0 1-.846.868C15.96 18.413 14.743 19 13 19c-1.743 0-2.96-.587-3.75-1.22a4.838 4.838 0 0 1-.847-.867a3.816 3.816 0 0 1-.286-.443l-.006-.013l-.003-.005l-.001-.003c-.001-.001-.001-.002.893-.449l-.894.447a1 1 0 1 1 1.798-.88c.017.03.05.08.1.146c.098.13.26.317.496.506c.46.368 1.243.781 2.5.781s2.04-.413 2.5-.78a2.84 2.84 0 0 0 .497-.507a1.826 1.826 0 0 0 .114-.172a1 1 0 0 1 1.336-.436Z" clip-rule="evenodd"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopFaceSmilingCircleFilled0)"/></g>`),
		g.Group(children),
	)
}