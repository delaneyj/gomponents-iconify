package bi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrightnessAltLow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8.5 5.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0zm5 6a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1zM2 11a.5.5 0 1 0 1 0a.5.5 0 0 0-1 0zm10.243-3.536a.5.5 0 1 1-.707-.707a.5.5 0 0 1 .707.707zm-8.486-.707a.5.5 0 1 0 .707.707a.5.5 0 0 0-.707-.707zM8 7a4 4 0 0 0-4 4a.5.5 0 0 0 .5.5h7a.5.5 0 0 0 .5-.5a4 4 0 0 0-4-4zm0 1a3 3 0 0 1 2.959 2.5H5.04A3 3 0 0 1 8 8z"/>`),
		g.Group(children),
	)
}