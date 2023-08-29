package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Layerorder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M985 559L607 748q-40 20-95 20t-95-20L39 559Q0 539 0 511.5T39 465l98-49l280 140q40 20 95 20t95-20l280-140l98 49q39 19 39 47t-39 47zm0-256L607 492q-40 20-95 20t-95-20L39 303Q0 283 0 255.5T39 208L417 19q40-19 95-19t95 19l378 189q39 20 39 47.5T985 303zM512 64L64 256l448 192l448-192zM39 721l98-49l70 35l-143 61l448 192l448-192l-143-61l70-35l98 49q39 19 39 47t-39 47l-378 189q-40 20-95 20t-95-20L39 815Q0 796 0 768t39-47z"/>`),
		g.Group(children),
	)
}