package websymbol

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Folder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1000 1000"),
		g.Raw(`<path fill="currentColor" d="M1313 218H78q9-48 27-67q16-17 31-17h4q4 1 8 1q12 0 26-4q18-6 25-29l24-72Q300 1 431 1q89 0 164 29l25 72q24 23 32.5 26t43.5 3h485q84 0 110 22q12 11 22 65zm78 164v10l-68 523q-5 35-33.5 60.5T1225 1001H167q-36 0-65.5-25.5T68 915L1 392q-1-3-1-10q0-33 22.5-54.5T78 306h1235q33 0 55.5 21.5T1391 382z"/>`),
		g.Group(children),
	)
}