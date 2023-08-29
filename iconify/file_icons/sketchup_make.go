package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SketchupMake(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m313.775 0l160.987 125.144l-34.227 292.705L109.843 512L46.532 314.085L0 40.3L313.775 0zM52.91 108.812L249.32 80.46l70.58 72.669l115.334-20.452L305.157 31.553L34.94 66.329l17.97 42.483zm13.341 82.781l18.276 48.043l99.705-19.107l25.985 44.546l108.67-23.252l-67.227-82.792l-185.41 32.562zm31.986 138.13l17.965 51.9l100.127-25.086l-24.546-48.225l-93.546 21.41z"/>`),
		g.Group(children),
	)
}