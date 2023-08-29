package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Off(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-139 0-257-68.5T68.5 769T0 512q0-97 34.5-185T132 170q3-3 11.5-12t14-14.5t15-10.5t19.5-5q26 0 45 19t19 45q0 16-3.5 28t-7.5 16l-3 5q-54 53-84 123t-30 148q0 104 51.5 192.5t140 140T512 896t192.5-51.5t140-140T896 512q0-78-30-148.5T783 240q-15-15-15-48q0-26 19-45t45-19q7 0 13.5 2t10.5 4t10.5 8t8.5 8.5t9 10.5t8 9q63 69 97.5 157t34.5 185q0 104-40.5 199t-109 163.5t-163.5 109t-199 40.5zm0-512q-27 0-45.5-18.5T448 448V64q0-27 19-45.5T512.5 0t45 18.5T576 64v384q0 27-18.5 45.5T512 512z"/>`),
		g.Group(children),
	)
}