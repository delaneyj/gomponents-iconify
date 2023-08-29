package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Delete(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M237 0h-42q-28 0-46 18.5T131 64v21H45q-17 0-29.5 13T3 128v21q0 18 12.5 30.5T45 192h2l39 284q1 16 13.5 26t29.5 10h174q17 0 29.5-10t13.5-26l39-284h2q17 0 29.5-12.5T429 149v-21q0-17-12.5-30T387 85h-86V64q0-27-18-45.5T237 0zm-64 64q0-21 22-21h42q22 0 22 21v21h-86V64zm-44 405l-3-21h182l-2 21H129zm183-64H120L90 192h250zm75-277v21H45v-21h342z"/>`),
		g.Group(children),
	)
}