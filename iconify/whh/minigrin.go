package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Minigrin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1024 513q0 73-35.5 141.5T889 777t-162.5 87T512 897t-214.5-33T135 777T35.5 654T0 513q0-27 19-45.5T64 449h896q26 0 45 18.5t19 45.5zm-880 64q17 35 47.5 68t75 61.5T375 752t137 17q74 0 138-17t108.5-45.5t75-61T881 577H144zm738.5-271q-13.5 15-33 15T816 306l-145-97q-11-4-18-12q-14-15-14-36.5t14-36.5q7-7 18-11l145-98q14-15 33.5-15t33 15T896 51.5T882 88l-108 73l108 72q14 15 14 36.5T882.5 306zM353 209l-145 97q-14 15-33.5 15t-33-15t-13.5-36.5t14-36.5l108-72l-108-73q-14-15-14-36.5T141.5 15t33-15T208 15l145 98q11 4 18 11q14 15 14 36.5T371 197q-7 8-18 12z"/>`),
		g.Group(children),
	)
}