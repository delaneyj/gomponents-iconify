package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M176 341q45 0 76-31t31-75V107q0-45-31-76T176 0t-76 31t-31 76v128q0 44 31 75t76 31zm-64-234q0-28 18.5-46T176 43t45.5 18t18.5 46v128q0 27-18.5 45.5T176 299t-45.5-18.5T112 235V107zm235 128v-64q0-22-22-22q-9 0-15 6t-6 16v64q0 53-38 90.5T176 363t-90-37.5T48 235v-64q0-10-6-16t-15-6q-22 0-22 22v64q0 65 43 112.5T155 403v45q0 7-5.5 12t-11.5 7l-5 2q-17 0-29.5 13T91 512h170q0-17-12.5-30T219 469q-22-6-22-21v-45q64-8 107-55.5T347 235z"/>`),
		g.Group(children),
	)
}