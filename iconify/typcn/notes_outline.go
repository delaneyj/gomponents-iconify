package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotesOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.324 4.367c-.368-.324-.84-.5-1.324-.5l-.248.016l-9 1.25A1.999 1.999 0 0 0 6 7.117v6.111c-1.746.551-3 2.034-3 3.772c0 2.205 2.019 4 4.5 4c1.695 0 3.169-.842 3.937-2.078A4.788 4.788 0 0 0 14.5 20c2.481 0 4.5-1.795 4.5-4V5.867c0-.574-.246-1.119-.676-1.5zM11 16v-4.256l3-.45v1.737c-1.693.208-3 1.46-3 2.969zm6 0c0 1.104-1.119 2-2.5 2s-2.5-.896-2.5-2s1.119-2 2.5-2c.172 0 .338.014.5.041v-3.908l-5 .75V17c0 1.104-1.119 2-2.5 2S5 18.104 5 17s1.119-2 2.5-2c.172 0 .338.014.5.041V7.117l9-1.25V16z"/>`),
		g.Group(children),
	)
}