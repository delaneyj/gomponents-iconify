package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeDownSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.977 4.201a1 1 0 0 1 1.633.645l.095.766c.53 4.242.53 8.534 0 12.776l-.095.766a1 1 0 0 1-1.633.644l-5.019-4.182a.5.5 0 0 0-.32-.116H4.75c-.69 0-1.25-.56-1.25-1.25v-4.5c0-.69.56-1.25 1.25-1.25h3.888a.5.5 0 0 0 .32-.116l5.02-4.183Zm5.055 3.747a.75.75 0 0 0-1.408.517c.405 1.1.626 2.291.626 3.535a10.25 10.25 0 0 1-.626 3.535a.75.75 0 0 0 1.408.517A11.667 11.667 0 0 0 19.75 12c0-1.423-.253-2.788-.718-4.052Z"/>`),
		g.Group(children),
	)
}