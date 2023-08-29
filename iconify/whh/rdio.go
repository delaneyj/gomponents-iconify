package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rdio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M889 371q7 41 7 77q0 91-35.5 174T765 765t-143 95.5T448 896t-174-35.5T131 765T35.5 622T0 448t35.5-174T131 131t143-95.5T448 0q63 0 128 19v259q-38-22-96-22q-52 0-104.5 33.5t-86 86T256 480q0 160 160 160q52 0 104.5-33.5t86-86T640 416V44q99 47 165 134q117 78 219 78q0 34-39 66t-96 49z"/>`),
		g.Group(children),
	)
}