package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MentionSquareBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M12 10.35a1.65 1.65 0 1 1 0 3.3a1.65 1.65 0 0 1 0-3.3Z"/><path fill-rule="evenodd" d="M3.464 3.464C2 4.93 2 7.286 2 12c0 4.714 0 7.071 1.464 8.535C4.93 22 7.286 22 12 22c4.714 0 7.071 0 8.535-1.465C22 19.072 22 16.714 22 12s0-7.071-1.465-8.536C19.072 2 16.714 2 12 2S4.929 2 3.464 3.464ZM6.75 12a5.25 5.25 0 1 1 10.184 1.797a.796.796 0 0 1-.216.304l-.08.075a.887.887 0 0 1-1.488-.652V12a3.15 3.15 0 1 0-1.277 2.533a2.386 2.386 0 0 0 3.782.745l.08-.074c.25-.23.476-.528.609-.893A6.75 6.75 0 1 0 12 18.75a.75.75 0 0 0 0-1.501A5.25 5.25 0 0 1 6.75 12Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}