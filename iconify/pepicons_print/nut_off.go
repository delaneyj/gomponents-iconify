package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NutOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><g opacity=".2"><path d="M7.58 4h6.84a1 1 0 0 1 .868.504l3.428 6a1 1 0 0 1 0 .992l-3.428 6a1 1 0 0 1-.868.504H7.58a1 1 0 0 1-.868-.504l-3.428-6a1 1 0 0 1 0-.992l3.428-6A1 1 0 0 1 7.58 4Z"/><path fill-rule="evenodd" d="m8.16 6l-2.857 5l2.858 5h5.678l2.857-5l-2.857-5H8.161Zm6.26-2H7.58a1 1 0 0 0-.868.504l-3.428 6a1 1 0 0 0 0 .992l3.428 6A1 1 0 0 0 7.58 18h6.84a1 1 0 0 0 .868-.504l3.428-6a1 1 0 0 0 0-.992l-3.428-6A1 1 0 0 0 14.42 4Z" clip-rule="evenodd"/></g><path fill-rule="evenodd" d="M14.722 16.744a1.5 1.5 0 0 1-1.302.756H6.58a1.5 1.5 0 0 1-1.302-.756l-3.429-6a1.5 1.5 0 0 1 0-1.488l3.429-6A1.5 1.5 0 0 1 6.58 2.5h6.84a1.5 1.5 0 0 1 1.302.756l3.429 6a1.5 1.5 0 0 1 0 1.488l-3.429 6ZM13.42 16.5a.5.5 0 0 0 .434-.252l3.428-6a.5.5 0 0 0 0-.496l-3.428-6a.5.5 0 0 0-.434-.252H6.58a.5.5 0 0 0-.434.252l-3.428 6a.5.5 0 0 0 0 .496l3.428 6a.5.5 0 0 0 .434.252h6.84Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M10 7.5a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5ZM6.5 10a3.5 3.5 0 1 1 7 0a3.5 3.5 0 0 1-7 0Z" clip-rule="evenodd"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}