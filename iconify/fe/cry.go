package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cry(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feCry0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feCry1" fill="currentColor"><path id="feCry2" d="M9.5 12a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm5.5 4a3 3 0 0 0-6 0h1s.317-2 2-2s1.996 2 1.996 2H15Zm-.5-4a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3ZM12 20a8 8 0 1 0 0-16a8 8 0 0 0 0 16Zm0 2C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10Zm4-7a1 1 0 0 0 1-1c0-.368-.333-1.035-1-2c-.667.965-1 1.632-1 2a1 1 0 0 0 1 1Z"/></g></g>`),
		g.Group(children),
	)
}