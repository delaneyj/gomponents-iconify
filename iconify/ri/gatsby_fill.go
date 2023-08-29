package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GatsbyFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.001 2c-5.5 0-10 4.5-10 10s4.5 10 10 10s10-4.5 10-10s-4.5-10-10-10ZM6.43 17.571c-1.5-1.5-2.286-3.5-2.286-5.428l7.786 7.714c-2-.071-4-.786-5.5-2.286Zm7.285 2.072l-9.357-9.357c.786-3.5 3.929-6.143 7.643-6.143c2.643 0 4.929 1.286 6.357 3.214l-1.071.929C16.072 6.643 14.144 5.57 12 5.57c-2.786 0-5.143 1.786-6.071 4.286l8.214 8.214c2.071-.714 3.643-2.5 4.143-4.642h-3.429V12h5c0 3.714-2.643 6.857-6.143 7.643Z"/>`),
		g.Group(children),
	)
}