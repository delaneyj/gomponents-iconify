package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StrolleyPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><g transform="translate(1)"><ellipse cx="4.438" cy="14.459" rx="1.438" ry="1.459"/><ellipse cx="12.457" cy="14.419" rx="1.457" ry="1.419"/><path d="M15 10.016H3.953V3.778L1.781.41A.995.995 0 1 0 .174 1.584l1.849 2.901v5.531h-.055c-.553 0-1 .44-1 .984c0 .543.447.984 1 .984H15c.553 0 1-.441 1-.984a.991.991 0 0 0-1-.984z"/></g><path d="M12.959 4.025H9.063c-.558 0-1.01.434-1.01.969c0 .535.452.969 1.01.969h3.896c.559 0 1.01-.434 1.01-.969c0-.535-.451-.969-1.01-.969z"/><path d="M11.979 6.942V3.046c0-.558-.434-1.01-.969-1.01c-.535 0-.969.452-.969 1.01v3.896c0 .559.434 1.01.969 1.01c.535 0 .969-.451.969-1.01z"/></g>`),
		g.Group(children),
	)
}