package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EyeCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13 19.5c4.897 0 9-2.308 9-5.5s-4.103-5.5-9-5.5s-9 2.308-9 5.5s4.103 5.5 9 5.5Zm0-9c3.94 0 7 1.722 7 3.5s-3.06 3.5-7 3.5s-7-1.722-7-3.5s3.06-3.5 7-3.5Z" clip-rule="evenodd"/><path d="M12 6.5a1 1 0 0 1 2 0v3a1 1 0 1 1-2 0v-3Zm4.02.304a1 1 0 0 1 1.96.392l-.5 2.5a1 1 0 0 1-1.96-.392l.5-2.5Zm-6.04 0a1 1 0 0 0-1.96.392l.5 2.5a1 1 0 0 0 1.96-.392l-.5-2.5ZM5.857 7.986a1 1 0 1 0-1.714 1.029l1.5 2.5a1 1 0 1 0 1.714-1.03l-1.5-2.5Zm14.285 0a1 1 0 0 1 1.716 1.029l-1.5 2.5a1 1 0 0 1-1.716-1.03l1.5-2.5Z"/><path fill-rule="evenodd" d="M13 17a3.5 3.5 0 1 0 0-7a3.5 3.5 0 0 0 0 7Zm0-5a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Z" clip-rule="evenodd"/><path d="M4.293 5.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/><path fill-rule="evenodd" d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}