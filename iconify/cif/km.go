package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Km(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#FFC61E" d="M.5.5h300v180H.5z"/><path fill="#FFF" d="M.5 45.5h300v135H.5z"/><path fill="#CE1126" d="M.5 90.5h300v90H.5z"/><path fill="#3A75C4" d="M.5 135.5h300v45H.5z"/><path fill="#3D8E33" d="m.5 180.5l150-90L.5.5z"/><circle cx="51.5" cy="90.5" r="40.5" fill="#FFF"/><circle cx="69.5" cy="90.5" r="40.5" fill="#3D8E33"/><path fill="#FFF" d="M53.368 59.195h5.449l1.683-5.183l1.683 5.183h5.449l-4.408 3.202l1.683 5.182l-4.407-3.202l-4.408 3.202l1.683-5.182zm0 19.325h5.449l1.683-5.182l1.683 5.182h5.449l-4.408 3.202l1.683 5.183l-4.407-3.202l-4.408 3.202l1.683-5.183zm0 19.325h5.449l1.683-5.183l1.683 5.183h5.449l-4.408 3.202l1.683 5.183l-4.407-3.203l-4.408 3.203l1.683-5.183zm0 19.325h5.449l1.683-5.183l1.683 5.183h5.449l-4.408 3.202l1.683 5.183l-4.407-3.202l-4.408 3.202l1.683-5.183z"/></g>`),
		g.Group(children),
	)
}