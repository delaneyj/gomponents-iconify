package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForPakistan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#e6e7e8" d="M10 10C3.373 10 0 14.925 0 21v22c0 6.075 3.373 11 10 11h10V10H10z"/><path fill="#006838" d="M54 10H20v44h34c6.627 0 10-4.925 10-11V21c0-6.075-3.373-11-10-11"/><g fill="#e6e7e8"><path d="m44.509 21.528l2.264 2.101l3.01-.671l-1.298 2.806l1.57 2.66l-3.073-.374l-2.044 2.319l-.593-3.03l-2.84-1.231l2.7-1.506z"/><path d="M52.37 35.33c-4.301 4.302-11.272 4.302-15.574 0c-4.302-4.301-4.302-11.273 0-15.574c.297-.296.607-.567.927-.822a12.638 12.638 0 0 0-5.829 3.312c-4.966 4.967-4.966 13.02 0 17.987c4.969 4.966 13.02 4.966 17.987 0a12.638 12.638 0 0 0 3.312-5.829a11.48 11.48 0 0 1-.823.926"/></g>`),
		g.Group(children),
	)
}