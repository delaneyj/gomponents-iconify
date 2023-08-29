package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RewindCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilRewindCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M19.5 17.796L14.019 13L19.5 8.204v9.592Zm-6.14-4.043a1 1 0 0 1 0-1.506l5.482-4.796c.646-.566 1.658-.106 1.658.753v9.592a1 1 0 0 1-1.659.753l-5.48-4.796Z"/><path d="M12.5 17.796L7.019 13L12.5 8.204v9.592Zm-6.14-4.043a1 1 0 0 1 0-1.506l5.481-4.796c.647-.566 1.659-.106 1.659.753v9.592a1 1 0 0 1-1.659.753l-5.48-4.796Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilRewindCircleFilled0)"/></g>`),
		g.Group(children),
	)
}