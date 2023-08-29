package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackEh0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackEh1)"><use href="#flagpackEh0"/><path fill="#F7FCFF" fill-rule="evenodd" d="M0 0h32v24H0V0Z" clip-rule="evenodd"/><path fill="#272727" fill-rule="evenodd" d="M0 0v8h32V0H0Z" clip-rule="evenodd"/><path fill="#5EAA22" fill-rule="evenodd" d="M0 16v8h32v-8H0Z" clip-rule="evenodd"/><path fill="#E31D1C" fill-rule="evenodd" d="m0 0l16 12L0 24V0Zm21.688 15.477s-2.275-1.375-2.275-3.646c0-2.272 2.275-3.467 2.275-3.467c-1.018-.646-4.549.276-4.549 3.54c0 3.265 3.498 3.869 4.549 3.573Zm1.84-4.165l-1.323-1.24l.457 1.584l-1.281.607l1.573.512l.683 1.714l.3-1.587l1.553.343l-1.175-1.148l.397-1.225l-1.184.44Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackEh1"><use href="#flagpackEh0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}