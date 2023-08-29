package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Parkingmeter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m740 787l-54 183q-3 23-23.5 38.5T617 1024H279q-25 0-45.5-15.5T210 970l-54-183Q83 724 41.5 636T0 448q0-91 35.5-174T131 131t143-95.5T448 0t174 35.5T765 131t95.5 143T896 448q0 100-41.5 188T740 787zM288 640h320q13 0 22.5-9.5T640 608t-9.5-22.5T608 576H288q-13 0-22.5 9.5T256 608t9.5 22.5T288 640zM448 64q-143 0-252 91.5T64 384q152 56 337 63l181-182q16-15 36-7q-73-66-170-66q-82 0-147.5 46.5T208 360q-35-7-63-15q33-96 116-156.5T448 128t187 60.5T751 345q-28 8-63 15q-16-45-51-83q8 20-7 36L495 447q185-7 337-63q-22-137-131.5-228.5T448 64z"/>`),
		g.Group(children),
	)
}