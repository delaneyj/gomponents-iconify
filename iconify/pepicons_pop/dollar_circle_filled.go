package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DollarCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopDollarCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M10.897 8.7c-.551.413-.8.908-.8 1.37c0 .463.249.958.8 1.372c.552.414 1.36.7 2.295.7a1 1 0 1 1 0 2c-1.326 0-2.565-.402-3.495-1.1c-.93-.698-1.6-1.738-1.6-2.971c0-1.234.67-2.274 1.6-2.972c.93-.697 2.17-1.099 3.495-1.099c2.053 0 3.994.983 4.766 2.62a1 1 0 0 1-1.81.853C15.798 8.726 14.706 8 13.193 8c-.935 0-1.743.286-2.295.7Z"/><path d="M15.157 17.583c.551-.413.799-.908.799-1.37c0-.464-.248-.959-.8-1.372c-.551-.414-1.36-.7-2.294-.7a1 1 0 1 1 0-2c1.326 0 2.565.402 3.495 1.1c.93.698 1.599 1.738 1.599 2.971s-.669 2.274-1.6 2.971c-.93.698-2.168 1.1-3.494 1.1c-2.053 0-3.995-.983-4.766-2.621a1 1 0 0 1 1.809-.853c.352.748 1.444 1.474 2.957 1.474c.935 0 1.743-.286 2.295-.7ZM13 4a1 1 0 0 1 1 1v1a1 1 0 1 1-2 0V5a1 1 0 0 1 1-1Z"/><path d="M13 19a1 1 0 0 1 1 1v1a1 1 0 1 1-2 0v-1a1 1 0 0 1 1-1Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopDollarCircleFilled0)"/></g>`),
		g.Group(children),
	)
}