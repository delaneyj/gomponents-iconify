package bi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrightnessLowFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M12 8a4 4 0 1 1-8 0a4 4 0 0 1 8 0zM8.5 2.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0zm0 11a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0zm5-5a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1zm-11 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1zm9.743-4.036a.5.5 0 1 1-.707-.707a.5.5 0 0 1 .707.707zm-7.779 7.779a.5.5 0 1 1-.707-.707a.5.5 0 0 1 .707.707zm7.072 0a.5.5 0 1 1 .707-.707a.5.5 0 0 1-.707.707zM3.757 4.464a.5.5 0 1 1 .707-.707a.5.5 0 0 1-.707.707z"/>`),
		g.Group(children),
	)
}