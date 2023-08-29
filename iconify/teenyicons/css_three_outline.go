package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CssThreeOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M.5.5V0a.5.5 0 0 0-.498.542L.5.5Zm14 0l.498.042A.5.5 0 0 0 14.5 0v.5Zm-1 12l.158.474a.5.5 0 0 0 .34-.433L13.5 12.5Zm-6 2l-.158.474a.499.499 0 0 0 .316 0L7.5 14.5Zm-6-2l-.498.041a.5.5 0 0 0 .34.433L1.5 12.5Zm9-9h.5V3h-.5v.5Zm0 6l.158.474L11 9.86V9.5h-.5Zm-3 1l-.158.474l.158.053l.158-.053L7.5 10.5Zm-3-1H4v.36l.342.114L4.5 9.5ZM.5 1h14V0H.5v1ZM14.002.458l-1 12l.996.083l1-12l-.996-.083Zm-.66 11.568l-6 2l.316.948l6-2l-.316-.948Zm-5.684 2l-6-2l-.316.948l6 2l.316-.948Zm-5.66-1.567l-1-12l-.996.083l1 12l.996-.083ZM10.5 3H4v1h6.5V3ZM6 7h4.5V6H6v1Zm4-.5v3h1v-3h-1Zm.342 2.526l-3 1l.316.948l3-1l-.316-.948Zm-2.684 1l-3-1l-.316.948l3 1l.316-.948ZM5 9.5V8H4v1.5h1Zm5-6v3h1v-3h-1Z"/>`),
		g.Group(children),
	)
}