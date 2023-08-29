package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circlenextempty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zm0-960q-91 0-174 35.5T195 195T99.5 338T64 512t35.5 174T195 829t143 95.5T512 960t174-35.5T829 829t95.5-143T960 512t-35.5-174T829 195T686 99.5T512 64zm313 468L607 763q-12 13-31-7V268q18-20 31-6l218 231q7 8 7 19.5t-7 19.5zM384 768h-64q-26 0-45-18.5T256 704V320q0-26 19-45t45-19h64q27 0 45.5 19t18.5 45v384q0 27-18.5 45.5T384 768z"/>`),
		g.Group(children),
	)
}