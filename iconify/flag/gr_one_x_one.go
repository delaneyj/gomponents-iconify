package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GrOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#0d5eaf" fill-rule="evenodd" d="M0 0h512v57H0z"/><path fill="#fff" fill-rule="evenodd" d="M0 57h512v57H0z"/><path fill="#0d5eaf" fill-rule="evenodd" d="M0 114h512v57H0z"/><path fill="#fff" fill-rule="evenodd" d="M0 171h512v57H0z"/><path fill="#0d5eaf" fill-rule="evenodd" d="M0 228h512v56.9H0z"/><path fill="#fff" fill-rule="evenodd" d="M0 284.9h512v57H0z"/><path fill="#0d5eaf" fill-rule="evenodd" d="M0 341.9h512v57H0z"/><path fill="#fff" fill-rule="evenodd" d="M0 398.9h512v57H0z"/><path fill="#0d5eaf" d="M0 0h284.9v284.9H0z"/><g fill="#fff" fill-rule="evenodd" stroke-width="1.3"><path d="M114 0h57v284.9h-57z"/><path d="M0 114h284.9v57H0z"/></g><path fill="#0d5eaf" fill-rule="evenodd" d="M0 455h512v57H0z"/>`),
		g.Group(children),
	)
}