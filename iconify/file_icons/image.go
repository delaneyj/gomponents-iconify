package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Image(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M329.143 0H0v512h438.857V109.714L329.143 0zm73.143 475.429H36.57V36.57h256l109.715 109.715v329.143zM219.429 329.143c40.357 0 73.142 32.786 73.142 73.143h73.143V182.857l-73.143-73.143H73.143v292.572h73.143c0-40.357 32.768-73.143 73.143-73.143c-56.097 0-91.348-51.145-63.3-99.797s98.55-48.652 126.6 0s-7.203 99.797-63.3 99.797z"/>`),
		g.Group(children),
	)
}