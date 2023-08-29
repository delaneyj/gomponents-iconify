package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Layers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M985 559L607 747q-40 20-95 20t-95-20L39 559Q0 539 0 511.5T39 464l34-17l344 173q40 19 95 19t95-19l344-173l34 17q39 20 39 47.5T985 559zm0-256L607 492q-40 19-95 19t-95-19L39 303Q0 283 0 255.5T39 208L417 19q40-19 95-19t95 19l378 189q39 20 39 47.5T985 303zM39 720l34-17l344 172q40 20 95 20t95-20l344-172l34 17q39 20 39 47.5T985 815l-378 188q-40 20-95 20t-95-20L39 815Q0 795 0 767.5T39 720z"/>`),
		g.Group(children),
	)
}