package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Teapot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 320q0 128-64 128q-3 0-5 5q5 32 5 59q0 108-54.5 214.5T699 896H325q-99-71-155-197q-18 5-42 5q-34 0-64-46.5T17 545T0 416q0-160 128-160q51 0 80 22q54-70 133.5-110T512 128q97 0 181 45.5T830 297q2-6 2-9q0-66 47-113t113-47q16 0 32 3v125q-27 0-45.5 18.5T960 320zm-832 0q-64 0-64 114q0 57 27.5 120.5T147 635q-19-65-19-123q0-85 37-163q-17-29-37-29zM448 63.5q0-26.5 18.5-45t45-18.5T557 18.5t19 45t-19 45.5t-45.5 19t-45-19T448 63.5z"/>`),
		g.Group(children),
	)
}