package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Barrel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M800.857 640q13 0 22.5 9.5t9.5 22.5t-9.5 22.5t-22.5 9.5h-32v256h32q13 0 22.5 9.5t9.5 22.5t-9.5 22.5t-22.5 9.5h-768q-13 0-22.5-9.5T.857 992t9.5-22.5t22.5-9.5h32V704h-32q-13 0-22.5-9.5T.857 672t9.5-22.5t22.5-9.5h32V384h-32q-13 0-22.5-9.5T.857 352t9.5-22.5t22.5-9.5h32V64h-32q-13 0-22.5-9.5T.857 32t9.5-22.5t22.5-9.5h768q13 0 22.5 9.5t9.5 22.5t-9.5 22.5t-22.5 9.5h-32v256h32q13 0 22.5 9.5t9.5 22.5t-9.5 22.5t-22.5 9.5h-32v256h32zm-271-233l-113-151l-113 151q-47 51-47 123t47 123t113 51t113-51t47-123t-47-123z"/>`),
		g.Group(children),
	)
}