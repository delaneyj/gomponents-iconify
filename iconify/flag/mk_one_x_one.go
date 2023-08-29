package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MkOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#d20000" d="M0 0h512v512H0z"/><path fill="#ffe600" d="M0 0h86.8L256 246.9L425.2 0H512L0 512h86.8L256 265.1L425.2 512H512zm512 204.8v102.4L0 204.8v102.4zM204.8 0L256 219.4L307.2 0zm0 512L256 292.6L307.2 512z"/><circle cx="256" cy="256" r="82.3" fill="#ffe600" stroke="#d20000" stroke-width="18.3"/>`),
		g.Group(children),
	)
}