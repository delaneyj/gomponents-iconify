package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dentistry(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 2.125q1.65 0 2.825 1.175T21 6.125q0 .275-.038.738t-.112 1.062L19.475 18q-.125.95-.863 1.55t-1.687.6q-.575 0-1.063-.25t-.812-.7l-2.675-3.9q-.05-.1-.163-.138t-.237-.037q-.1 0-.4.225l-2.6 3.775q-.35.5-.863.763t-1.087.262q-.95 0-1.675-.612t-.85-1.563L3.15 7.925q-.075-.6-.113-1.062T3 6.125Q3 4.475 4.175 3.3T7 2.125q.9 0 1.438.238t1.037.512q.5.275 1.063.513T12 3.625q.9 0 1.463-.238t1.062-.512q.5-.275 1.05-.513T17 2.125Z"/>`),
		g.Group(children),
	)
}