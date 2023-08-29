package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Terminal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M475.429 36.571H36.57C16.371 36.571 0 52.943 0 73.143v365.714c0 20.2 16.371 36.572 36.571 36.572H475.43c20.2 0 36.571-16.372 36.571-36.572V73.143c0-20.2-16.371-36.572-36.571-36.572zm-402.286 256l73.143-73.142l-73.143-73.143l36.571-36.572L219.43 219.43L109.714 329.143L73.143 292.57zm292.571 36.572H219.43V292.57h146.285v36.572z"/>`),
		g.Group(children),
	)
}